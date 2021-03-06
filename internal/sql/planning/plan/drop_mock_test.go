// Code generated by MockGen. DO NOT EDIT.
// Source: drop.go

// Package plan_test is a generated GoMock package.
package plan_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatabaseDropper is a mock of DatabaseDropper interface
type MockDatabaseDropper struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseDropperMockRecorder
}

// MockDatabaseDropperMockRecorder is the mock recorder for MockDatabaseDropper
type MockDatabaseDropperMockRecorder struct {
	mock *MockDatabaseDropper
}

// NewMockDatabaseDropper creates a new mock instance
func NewMockDatabaseDropper(ctrl *gomock.Controller) *MockDatabaseDropper {
	mock := &MockDatabaseDropper{ctrl: ctrl}
	mock.recorder = &MockDatabaseDropperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseDropper) EXPECT() *MockDatabaseDropperMockRecorder {
	return m.recorder
}

// DropDatabase mocks base method
func (m *MockDatabaseDropper) DropDatabase(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropDatabase", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropDatabase indicates an expected call of DropDatabase
func (mr *MockDatabaseDropperMockRecorder) DropDatabase(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropDatabase", reflect.TypeOf((*MockDatabaseDropper)(nil).DropDatabase), name)
}

// MockTableDropper is a mock of TableDropper interface
type MockTableDropper struct {
	ctrl     *gomock.Controller
	recorder *MockTableDropperMockRecorder
}

// MockTableDropperMockRecorder is the mock recorder for MockTableDropper
type MockTableDropperMockRecorder struct {
	mock *MockTableDropper
}

// NewMockTableDropper creates a new mock instance
func NewMockTableDropper(ctrl *gomock.Controller) *MockTableDropper {
	mock := &MockTableDropper{ctrl: ctrl}
	mock.recorder = &MockTableDropperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTableDropper) EXPECT() *MockTableDropperMockRecorder {
	return m.recorder
}

// DropTable mocks base method
func (m *MockTableDropper) DropTable(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropTable", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropTable indicates an expected call of DropTable
func (mr *MockTableDropperMockRecorder) DropTable(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropTable", reflect.TypeOf((*MockTableDropper)(nil).DropTable), name)
}
