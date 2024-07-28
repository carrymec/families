package relationship

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github/carrymec/families/pkg"
	"go.uber.org/zap"
)

type Dao struct {
	lg            *zap.Logger
	sessionClient pkg.PersonSessionWithContext
}

// 保证Dao都实现了DaoClient接口
var _ DaoClient = (*Dao)(nil)

func NewRelationDao(lg *zap.Logger, sessionClient neo4j.SessionWithContext) *Dao {
	return &Dao{
		lg:            lg,
		sessionClient: sessionClient,
	}
}

type DaoClient interface {
	Create(ctx context.Context, relation Relationship) (int64, error)
	Query(ctx context.Context, query Query) ([]Relationship, error)
	Update(ctx context.Context, id int64, relation Relationship) error
	FindById(ctx context.Context, id int64) (Relationship, error)
	Delete(ctx context.Context, id int64) error
}

func (d *Dao) Create(ctx context.Context, relation Relationship) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Dao) Query(ctx context.Context, query Query) ([]Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Dao) Update(ctx context.Context, id int64, relation Relationship) error {
	//TODO implement me
	panic("implement me")
}

func (d *Dao) FindById(ctx context.Context, id int64) (Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Dao) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
