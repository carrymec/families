// Code generated by MockGen. DO NOT EDIT.
// Source: github/carrymec/families/person (interfaces: ServiceClient)

// Package mock_person is a generated GoMock package.
package mock_person

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	person "github/carrymec/families/person"
	reflect "reflect"
)

// MockServiceClient is a mock of ServiceClient interface
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// CheckExistByName mocks base method
func (m *MockServiceClient) CheckExistByName(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistByName", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistByName indicates an expected call of CheckExistByName
func (mr *MockServiceClientMockRecorder) CheckExistByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistByName", reflect.TypeOf((*MockServiceClient)(nil).CheckExistByName), arg0, arg1)
}

// CheckExistRelationship mocks base method
func (m *MockServiceClient) CheckExistRelationship(arg0 context.Context, arg1, arg2 int64, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistRelationship", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistRelationship indicates an expected call of CheckExistRelationship
func (mr *MockServiceClientMockRecorder) CheckExistRelationship(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistRelationship", reflect.TypeOf((*MockServiceClient)(nil).CheckExistRelationship), arg0, arg1, arg2, arg3)
}

// CreatePerson mocks base method
func (m *MockServiceClient) CreatePerson(arg0 context.Context, arg1 person.Person) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerson", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePerson indicates an expected call of CreatePerson
func (mr *MockServiceClientMockRecorder) CreatePerson(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerson", reflect.TypeOf((*MockServiceClient)(nil).CreatePerson), arg0, arg1)
}

// CreateRelationship mocks base method
func (m *MockServiceClient) CreateRelationship(arg0 context.Context, arg1, arg2 int64, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelationship", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRelationship indicates an expected call of CreateRelationship
func (mr *MockServiceClientMockRecorder) CreateRelationship(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelationship", reflect.TypeOf((*MockServiceClient)(nil).CreateRelationship), arg0, arg1, arg2, arg3)
}

// DeletePersonWithRelationship mocks base method
func (m *MockServiceClient) DeletePersonWithRelationship(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePersonWithRelationship", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePersonWithRelationship indicates an expected call of DeletePersonWithRelationship
func (mr *MockServiceClientMockRecorder) DeletePersonWithRelationship(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePersonWithRelationship", reflect.TypeOf((*MockServiceClient)(nil).DeletePersonWithRelationship), arg0, arg1, arg2)
}

// FindById mocks base method
func (m *MockServiceClient) FindById(arg0 context.Context, arg1 int64) (person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockServiceClientMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockServiceClient)(nil).FindById), arg0, arg1)
}

// Query mocks base method
func (m *MockServiceClient) Query(arg0 context.Context, arg1 person.Query) ([]person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].([]person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockServiceClientMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockServiceClient)(nil).Query), arg0, arg1)
}