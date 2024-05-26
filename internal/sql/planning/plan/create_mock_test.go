// Code generated by MockGen. DO NOT EDIT.
// Source: create.go
//
// Generated by this command:
//
//	mockgen -typed -source=create.go -destination ./create_mock_test.go -package plan_test
//

// Package plan_test is a generated GoMock package.
package plan_test

import (
	reflect "reflect"

	sql "github.com/i-sevostyanov/NanoDB/internal/sql"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabaseCreator is a mock of DatabaseCreator interface.
type MockDatabaseCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseCreatorMockRecorder
}

// MockDatabaseCreatorMockRecorder is the mock recorder for MockDatabaseCreator.
type MockDatabaseCreatorMockRecorder struct {
	mock *MockDatabaseCreator
}

// NewMockDatabaseCreator creates a new mock instance.
func NewMockDatabaseCreator(ctrl *gomock.Controller) *MockDatabaseCreator {
	mock := &MockDatabaseCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseCreator) EXPECT() *MockDatabaseCreatorMockRecorder {
	return m.recorder
}

// CreateDatabase mocks base method.
func (m *MockDatabaseCreator) CreateDatabase(name string) (sql.Database, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabase", name)
	ret0, _ := ret[0].(sql.Database)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabase indicates an expected call of CreateDatabase.
func (mr *MockDatabaseCreatorMockRecorder) CreateDatabase(name any) *MockDatabaseCreatorCreateDatabaseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabase", reflect.TypeOf((*MockDatabaseCreator)(nil).CreateDatabase), name)
	return &MockDatabaseCreatorCreateDatabaseCall{Call: call}
}

// MockDatabaseCreatorCreateDatabaseCall wrap *gomock.Call
type MockDatabaseCreatorCreateDatabaseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDatabaseCreatorCreateDatabaseCall) Return(arg0 sql.Database, arg1 error) *MockDatabaseCreatorCreateDatabaseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDatabaseCreatorCreateDatabaseCall) Do(f func(string) (sql.Database, error)) *MockDatabaseCreatorCreateDatabaseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDatabaseCreatorCreateDatabaseCall) DoAndReturn(f func(string) (sql.Database, error)) *MockDatabaseCreatorCreateDatabaseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockTableCreator is a mock of TableCreator interface.
type MockTableCreator struct {
	ctrl     *gomock.Controller
	recorder *MockTableCreatorMockRecorder
}

// MockTableCreatorMockRecorder is the mock recorder for MockTableCreator.
type MockTableCreatorMockRecorder struct {
	mock *MockTableCreator
}

// NewMockTableCreator creates a new mock instance.
func NewMockTableCreator(ctrl *gomock.Controller) *MockTableCreator {
	mock := &MockTableCreator{ctrl: ctrl}
	mock.recorder = &MockTableCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTableCreator) EXPECT() *MockTableCreatorMockRecorder {
	return m.recorder
}

// CreateTable mocks base method.
func (m *MockTableCreator) CreateTable(name string, scheme sql.Scheme) (sql.Table, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", name, scheme)
	ret0, _ := ret[0].(sql.Table)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockTableCreatorMockRecorder) CreateTable(name, scheme any) *MockTableCreatorCreateTableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockTableCreator)(nil).CreateTable), name, scheme)
	return &MockTableCreatorCreateTableCall{Call: call}
}

// MockTableCreatorCreateTableCall wrap *gomock.Call
type MockTableCreatorCreateTableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTableCreatorCreateTableCall) Return(arg0 sql.Table, arg1 error) *MockTableCreatorCreateTableCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTableCreatorCreateTableCall) Do(f func(string, sql.Scheme) (sql.Table, error)) *MockTableCreatorCreateTableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTableCreatorCreateTableCall) DoAndReturn(f func(string, sql.Scheme) (sql.Table, error)) *MockTableCreatorCreateTableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
