// Code generated by MockGen. DO NOT EDIT.
// Source: rows.go

// Package sql is a generated GoMock package.
package sql

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRowIter is a mock of RowIter interface
type MockRowIter struct {
	ctrl     *gomock.Controller
	recorder *MockRowIterMockRecorder
}

// MockRowIterMockRecorder is the mock recorder for MockRowIter
type MockRowIterMockRecorder struct {
	mock *MockRowIter
}

// NewMockRowIter creates a new mock instance
func NewMockRowIter(ctrl *gomock.Controller) *MockRowIter {
	mock := &MockRowIter{ctrl: ctrl}
	mock.recorder = &MockRowIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRowIter) EXPECT() *MockRowIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockRowIter) Next() (Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockRowIterMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRowIter)(nil).Next))
}

// Close mocks base method
func (m *MockRowIter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockRowIterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRowIter)(nil).Close))
}
