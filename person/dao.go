package person

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github/carrymec/families/common"
	"github/carrymec/families/pkg"
	"go.uber.org/zap"
)

type Dao struct {
	lg *zap.Logger
	//sessionClient  neo4j.SessionWithContext
	sessionClient pkg.PersonSessionWithContext
}

func NewPersonDao(lg *zap.Logger, sessionClient neo4j.SessionWithContext) *Dao {
	return &Dao{
		lg:            lg,
		sessionClient: sessionClient,
	}
}

type DaoClient interface {
	CreatePerson(ctx context.Context, person Person) (int64, error)
	CreateRelationship(ctx context.Context, fromId, toId int64, relationType common.RelationType) error
	CheckExistByName(ctx context.Context, name string) (bool, error)
	CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType common.RelationType) (bool, error)
	Query(ctx context.Context, query Query) ([]Person, error)
	Update(ctx context.Context, id int64, person Person) error
	FindById(ctx context.Context, id int64) (Person, error)
}

func (d *Dao) FindById(ctx context.Context, id int64) (Person, error) {
	per, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx,
			`MATCH (p:Person) WHERE id(p) = $id RETURN p.name as name, p.birthdate as birthdate, p.note as note`,
			map[string]any{
				"id": id,
			})
		if err != nil {
			d.lg.Error("search person by query err", zap.Int64("id", id), zap.Error(err))
			return nil, err
		}
		var p Person
		for res.Next(ctx) {
			record := res.Record()
			name, _ := record.Get("name")
			birthdate, _ := record.Get("birthdate")
			note, _ := record.Get("note")
			p = Person{
				ID:        id,
				Name:      name.(string),
				Birthdate: birthdate.(string),
				Note:      note.(string),
			}
			break
		}
		return p, nil
	})
	if err != nil {
		d.lg.Error("execute read err", zap.Error(err))
		return Person{}, err
	}
	return per.(Person), nil
}

func (d *Dao) Update(ctx context.Context, id int64, person Person) error {
	return nil
}

// 这里目前只查询了基础信息 后续可以带上关系
func (d *Dao) Query(ctx context.Context, query Query) ([]Person, error) {
	results, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := `MATCH (p:Person) RETURN id(p) as id, p.name as name, p.birthdate as birthdate SKIP $page LIMIT $pageSize`
		paramsMap := map[string]any{
			"page":     query.PageSize * (query.Page - 1), //pageSize * (page - 1)
			"pageSize": query.PageSize,
		}
		if query.Name != "" {
			cypher = `MATCH (p:Person) WHERE p.name CONTAINS $name RETURN id(p) as id, p.name as name, p.birthdate as birthdate SKIP $page LIMIT $pageSize`
			paramsMap["name"] = query.Name
		}
		res, err := tx.Run(ctx, cypher, paramsMap)
		if err != nil {
			d.lg.Error("search person by query err", zap.Any("query", query), zap.Error(err))
			return nil, err
		}
		var persons []Person
		for res.Next(ctx) {
			r := res.Record()
			id, _ := r.Get("id")
			name, _ := r.Get("name")
			birthdate, _ := r.Get("birthdate")
			note, _ := r.Get("note")
			persons = append(persons, Person{
				ID:        id.(int64),
				Name:      name.(string),
				Birthdate: birthdate.(string),
				Note:      note.(string),
			})
		}
		if err = res.Err(); err != nil {
			return nil, err
		}

		return persons, nil
	})
	if err != nil {
		d.lg.Error("execute read err", zap.Error(err))
		return nil, err
	}

	return results.([]Person), nil
}

func (d *Dao) CheckExistByName(ctx context.Context, name string) (bool, error) {
	userId, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `MATCH(p:Person{name: $name}) RETURN id(p) as id`,
			map[string]any{
				"name": name,
			})
		if err != nil {
			d.lg.Error("match person err", zap.Error(err))
			return nil, err
		}
		flag := false
		if res.Next(ctx) {
			a := res.Record().AsMap()["id"]
			idint := a.(int64)
			flag = idint != -1
			d.lg.Debug("match person ok", zap.Int64("id", idint))
		}
		return flag, nil
	})
	if err != nil {
		d.lg.Error("execute read err", zap.Error(err))
		return false, err
	}
	return userId.(bool), nil
}

/*
	创建并带上关系

MATCH (n:Person) WHERE id(n) = 427
create(p:Person{name: "秦王政1",birthdate: "前259年－前210年"})
CREATE (p)-[pson:son]->(n)
*/
func (d *Dao) CreatePerson(ctx context.Context, person Person) (int64, error) {
	id, err := d.sessionClient.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		cypher := "CREATE (p:Person {name: $name, birthdate: $birthdate, note: $note}) RETURN id(p) as id"
		paramMap := map[string]interface{}{
			"name":      person.Name,
			"birthdate": person.Birthdate,
			"note":      person.Note,
		}
		if person.Relation != nil {
			// 代表有关系绑定 关系ID在service层做了校验
			cypher = `MATCH(n:Person) WHERE id(n)=$relationId
                      CREATE(p:Person {name: $name, birthdate: $birthdate})
					  CREATE(p) -[r:` + string(person.Relation.RelationType) + `] ->(n)
                      RETURN id(p) as id`
			paramMap = map[string]interface{}{
				"relationId": person.Relation.RelationId,
				"name":       person.Name,
				"birthdate":  person.Birthdate,
				"note":       person.Note,
			}
		}

		result, err := tx.Run(
			ctx,
			cypher,
			paramMap,
		)
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

func (d *Dao) CreateRelationship(ctx context.Context, fromId, toId int64, relationType common.RelationType) error {
	_, err := d.sessionClient.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `MATCH (a:Person {id: $fromId}), (b:Person {id: $toId})
             CREATE (a)-[r:`+string(relationType)+`]->(b)
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

func (d *Dao) CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType common.RelationType) (bool, error) {
	exist, err := d.sessionClient.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		res, err := tx.Run(ctx, `
			 MATCH (a:Person {id: $fromId})-[r:`+string(relationType)+`]->(b:Person {id: $toId})
             RETURN COUNT(r) AS count`, map[string]any{
			"fromId": fromId,
			"toId":   toId,
		})
		if err != nil {
			d.lg.Error("match relationship err", zap.Int64("fromId", fromId), zap.Int64("toId", toId), zap.String("relationType", string(relationType)), zap.Error(err))
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
