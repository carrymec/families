// Code generated by MockGen. DO NOT EDIT.
// Source: github/carrymec/families/pkg (interfaces: ManagedTransaction)

// Package mock_session is a generated GoMock package.
package mock_session

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	neo4j "github.com/neo4j/neo4j-go-driver/v5/neo4j"
	reflect "reflect"
)

// MockManagedTransaction is a mock of ManagedTransaction interface
type MockManagedTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockManagedTransactionMockRecorder
}

// MockManagedTransactionMockRecorder is the mock recorder for MockManagedTransaction
type MockManagedTransactionMockRecorder struct {
	mock *MockManagedTransaction
}

// NewMockManagedTransaction creates a new mock instance
func NewMockManagedTransaction(ctrl *gomock.Controller) *MockManagedTransaction {
	mock := &MockManagedTransaction{ctrl: ctrl}
	mock.recorder = &MockManagedTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagedTransaction) EXPECT() *MockManagedTransactionMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockManagedTransaction) Run(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (neo4j.ResultWithContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2)
	ret0, _ := ret[0].(neo4j.ResultWithContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockManagedTransactionMockRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManagedTransaction)(nil).Run), arg0, arg1, arg2)
}

// legacy mocks base method
func (m *MockManagedTransaction) legacy() neo4j.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "legacy")
	ret0, _ := ret[0].(neo4j.Transaction)
	return ret0
}

// legacy indicates an expected call of legacy
func (mr *MockManagedTransactionMockRecorder) legacy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "legacy", reflect.TypeOf((*MockManagedTransaction)(nil).legacy))
}
