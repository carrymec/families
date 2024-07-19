package person

import (
	"context"
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
	CreatePerson(ctx context.Context, person Person) (int64, error)
	CreateRelationship(ctx context.Context, fromId, toId int64, relationType string) error
	CheckExistByName(ctx context.Context, name string) (int64, error)
	CheckExistRelationship(ctx context.Context, fromId, toId int64, relationType string) (bool, error)
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
