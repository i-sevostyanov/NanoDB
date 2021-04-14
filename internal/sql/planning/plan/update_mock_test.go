// Code generated by MockGen. DO NOT EDIT.
// Source: update.go

// Package plan_test is a generated GoMock package.
package plan_test

import (
	gomock "github.com/golang/mock/gomock"
	sql "github.com/i-sevostyanov/NanoDB/internal/sql"
	reflect "reflect"
)

// MockRowUpdater is a mock of RowUpdater interface
type MockRowUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockRowUpdaterMockRecorder
}

// MockRowUpdaterMockRecorder is the mock recorder for MockRowUpdater
type MockRowUpdaterMockRecorder struct {
	mock *MockRowUpdater
}

// NewMockRowUpdater creates a new mock instance
func NewMockRowUpdater(ctrl *gomock.Controller) *MockRowUpdater {
	mock := &MockRowUpdater{ctrl: ctrl}
	mock.recorder = &MockRowUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRowUpdater) EXPECT() *MockRowUpdaterMockRecorder {
	return m.recorder
}

// Update mocks base method
func (m *MockRowUpdater) Update(key int64, row sql.Row) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", key, row)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRowUpdaterMockRecorder) Update(key, row interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRowUpdater)(nil).Update), key, row)
}