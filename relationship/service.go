package relationship

import (
	"context"
	"go.uber.org/zap"
)

var _ ServiceClient = (*Service)(nil)

type Service struct {
	lg  *zap.Logger
	dao DaoClient
}

func NewService(lg *zap.Logger, client DaoClient) *Service {
	return &Service{
		lg:  lg,
		dao: client,
	}
}

type ServiceClient interface {
	// Create 创建用户
	Create(ctx context.Context, relation Relationship) (int64, error)
	// CheckExistByName 校验用户是否存在，这里通过name去校验，name唯一处理
	CheckExistByName(ctx context.Context, name string) (bool, error)
	// Query 条件查询并分页
	Query(ctx context.Context, query Query) ([]Relationship, error)
	// FindById 通过id查询
	FindById(ctx context.Context, id int64) (Relationship, error)
	// Update 通过id更新
	Update(ctx context.Context, id int64, relation Relationship) error
	// Delete 通过id删除
	Delete(ctx context.Context, id int64) error
}

func (s Service) Create(ctx context.Context, relation Relationship) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) CheckExistByName(ctx context.Context, name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Query(ctx context.Context, query Query) ([]Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) FindById(ctx context.Context, id int64) (Relationship, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(ctx context.Context, id int64, relation Relationship) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
