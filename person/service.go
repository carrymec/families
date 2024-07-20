package person

import (
	"context"
	"fmt"
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
}

func (s *Service) CreatePerson(ctx context.Context, person Person) (int64, error) {
	// 校验用户是否存在
	id, err := s.CheckExistByName(ctx, person.Name)
	if err != nil {
		s.lg.Error("校验用户名是否存在失败", zap.Error(err))
		return 0, err
	}
	if id != -1 {
		s.lg.Info("用户名已存在", zap.String("name", person.Name))
		return id, fmt.Errorf("用户名 %s 已存在", person.Name)
	}
	// 创建用户
	id, err = s.dao.CreatePerson(ctx, person)
	if err != nil {
		s.lg.Error("创建用户失败", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (s *Service) CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error {
	// 校验关系是否存在
	exist, err := s.CheckExistRelationship(ctx, fromId, toId, relationType)
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

func (s *Service) CheckExistByName(ctx context.Context, name string) (int64, error) {
	return s.dao.CheckExistByName(ctx, name)
}

func (s *Service) CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType string) (bool, error) {
	return s.dao.CheckExistRelationship(ctx, fromId, toId, relationType)
}

func (s *Service) DeletePersonWithRelationship(ctx context.Context, id int64, relationType string) error {
	return nil
}

/*
	personDao := person.NewPersonDao(logger.Logger, session)
	person1 := person.Person{ID: 1, Name: "John Doe", Birthdate: "1990-01-01"}
	person2 := person.Person{ID: 2, Name: "Jane Doe", Birthdate: "1992-02-02"}

	id1, err := personDao.CreatePerson(ctx, person1)
	if err != nil {
		panic(err)
	}
	id2, err := personDao.CreatePerson(ctx, person2)
	if err != nil {
		panic(err)
	}
	err = personDao.CreateRelationship(ctx, id1, id2, "PARENT_OF")
	if err != nil {
		panic(err)
	}
*/
