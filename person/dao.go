package person

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.uber.org/zap"
)

type PersonDao struct {
	lg            *zap.Logger
	sessionClient neo4j.SessionWithContext
}

func NewPersonDao(lg *zap.Logger, sessionClient neo4j.SessionWithContext) *PersonDao {
	return &PersonDao{
		lg:            lg,
		sessionClient: sessionClient,
	}
}

type DaoClient interface {
	CreatePerson(ctx context.Context, person Person) (int64, error)
	CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error
	CheckExistByName(ctx context.Context, name string) (int64, error)
	CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType string) (bool, error)
}

func (d *PersonDao) CheckExistByName(ctx context.Context, name string) (int64, error) {
	userId, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH(p:Person{name: $name}) RETURN p.id as id`, map[string]any{
			"name": name,
		})
		if err != nil {
			d.lg.Error("match person err", zap.Error(err))
			return nil, err
		}
		id := int64(-1)
		if res.Next(ctx) {
			a := res.Record().AsMap()["id"]
			id = a.(int64)
			d.lg.Debug("match person ok", zap.Int64("id", id))
		}
		return id, nil
	})
	if err != nil {
		d.lg.Error("execute read err", zap.Error(err))
		return -1, err
	}
	return userId.(int64), nil
}

func (d *PersonDao) CreatePerson(ctx context.Context, person Person) (int64, error) {
	id, err := d.sessionClient.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(
			ctx,
			"CREATE (p:Person {name: $name, birthdate: $birthdate}) RETURN id(p) as id",
			map[string]interface{}{
				"name":      person.Name,
				"birthdate": person.Birthdate,
			})
		if err != nil {
			d.lg.Error("create person err", zap.Error(err))
			return nil, err
		}

		per, err := result.Single(ctx)
		if err != nil {
			d.lg.Error("get result err", zap.Error(err))
			return "", err
		}

		id, _ := per.AsMap()["id"]

		return id, nil

	})
	if err != nil {
		d.lg.Error("execute write err", zap.Error(err))
		return 0, err
	}

	return id.(int64), nil
}

func (d *PersonDao) CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error {
	_, err := d.sessionClient.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `MATCH (a:Person {id: $fromId}), (b:Person {id: $toId})
             CREATE (a)-[r:`+relationType+`]->(b)
             RETURN type(r)`, map[string]any{
			"fromId": fromId,
			"toId":   toId,
		})
		if err != nil {
			d.lg.Error("create relationship err", zap.Error(err))
			return nil, err
		}

		if result.Next(ctx) {
			d.lg.Debug("Created Relationship", zap.Any("values", result.Record().Values))
		}
		return nil, nil
	})

	if err != nil {
		d.lg.Error("execute write err", zap.Error(err))
		return err
	}
	return nil
}

func (d *PersonDao) CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType string) (bool, error) {
	exist, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			 MATCH (a:Person {id: $fromId})-[r:`+relationType+`]->(b:Person {id: $toId})
             RETURN COUNT(r) AS count`, map[string]any{
			"fromId": fromId,
			"toId":   toId,
		})
		if err != nil {
			d.lg.Error("match relationship err", zap.Int64("fromId", fromId), zap.Int64("toId", toId), zap.String("relationType", relationType), zap.Error(err))
			return false, err
		}
		ok := int64(0)
		if res.Next(ctx) {
			a := res.Record().AsMap()["count"]
			ok = a.(int64)
		}
		return ok != 0, nil
	})
	if err != nil {
		d.lg.Error("execute read err", zap.Error(err))
		return false, err
	}
	return exist.(bool), nil
}
