package person

import (
	"context"
	"fmt"
	"github/carrymec/families/common"
	"go.uber.org/zap"
)

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
	// CreatePerson 创建用户
	CreatePerson(ctx context.Context, person Person) (int64, error)
	// CreateRelationship 创建关系
	CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error
	// CheckExistByName 校验用户是否存在，这里通过name去校验，name唯一处理
	CheckExistByName(ctx context.Context, name string) (int64, error)
	// CheckExistRelationship 校验关系是否存在
	CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType string) (bool, error)
	// DeletePersonWithRelationship 删除用户并删除关系
	DeletePersonWithRelationship(ctx context.Context, id int64, relationType string) error
	// Query 条件查询并分页
	Query(ctx context.Context, query Query) ([]Person, error)
	FindById(ctx context.Context, id int64) (Person, error)
}

func (s *Service) CreatePerson(ctx context.Context, person Person) (int64, error) {
	// 校验用户是否存在
	exist, err := s.CheckExistByName(ctx, person.Name)
	if err != nil {
		s.lg.Error("校验用户名是否存在失败", zap.Error(err))
		return 0, err
	}
	if exist {
		s.lg.Info("用户名已存在", zap.String("name", person.Name))
		return -1, fmt.Errorf("用户名 %s 已存在", person.Name)
	}
	// 创建用户
	id, err := s.dao.CreatePerson(ctx, person)
	if err != nil {
		s.lg.Error("创建用户失败", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (s *Service) CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error {
	// 校验关系是否存在
	exist, err := s.CheckExistRelationship(ctx, fromId, toId, common.RelationType(relationType))
	if err != nil {
		s.lg.Error("校验关系是否存在失败", zap.Error(err))
		return err
	}
	if exist {
		s.lg.Error("关系已存在", zap.Int64("fromId", fromId), zap.Int64("toId", toId), zap.String("relationType", relationType))
		return fmt.Errorf("关系已存在")
	}
	return nil
}

func (s *Service) CheckExistByName(ctx context.Context, name string) (bool, error) {
	return s.dao.CheckExistByName(ctx, name)
}

func (s *Service) CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType common.RelationType) (bool, error) {
	return s.dao.CheckExistRelationship(ctx, fromId, toId, relationType)
}

func (s *Service) DeletePersonWithRelationship(ctx context.Context, id int64, relationType string) error {
	return nil
}

func (s *Service) Query(ctx context.Context, query Query) ([]Person, error) {
	page := query.Page
	pageSize := query.PageSize
	if page == 0 && pageSize == 0 {
		query.Page = 1
		// 默认最大1000
		query.PageSize = 1000
	}
	peoples, err := s.dao.Query(ctx, query)
	if err != nil {
		s.lg.Error("query persons err", zap.Error(err))
		return nil, err
	}
	return peoples, nil
}

func (s *Service) FindById(ctx context.Context, id int64) (Person, error) {
	return s.dao.FindById(ctx, id)
}
