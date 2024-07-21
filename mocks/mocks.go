package mocks

//go:generate mockgen -package=mock_session -mock_names=SessionWithContext=MockSessionWithContext -destination=mock_session/mock_session.go github.com/neo4j/neo4j-go-driver/v5/neo4j SessionWithContext

//go:generate mockgen -package=mock_person -mock_names=Dao=MockPersonDao -destination=mock_person/mock_person_dao.go github/carrymec/families/person DaoClient

//go:generate mockgen -package=mock_person -mock_names=Service=MockPersonService -destination=mock_person/mock_person_service.go github/carrymec/families/person ServiceClient
