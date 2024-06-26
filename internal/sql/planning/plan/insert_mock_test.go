// Code generated by MockGen. DO NOT EDIT.
// Source: insert.go
//
// Generated by this command:
//
//	mockgen -typed -source=insert.go -destination ./insert_mock_test.go -package plan_test
//

// Package plan_test is a generated GoMock package.
package plan_test

import (
	reflect "reflect"

	sql "github.com/i-sevostyanov/NanoDB/internal/sql"
	gomock "go.uber.org/mock/gomock"
)

// MockTableInserter is a mock of TableInserter interface.
type MockTableInserter struct {
	ctrl     *gomock.Controller
	recorder *MockTableInserterMockRecorder
}

// MockTableInserterMockRecorder is the mock recorder for MockTableInserter.
type MockTableInserterMockRecorder struct {
	mock *MockTableInserter
}

// NewMockTableInserter creates a new mock instance.
func NewMockTableInserter(ctrl *gomock.Controller) *MockTableInserter {
	mock := &MockTableInserter{ctrl: ctrl}
	mock.recorder = &MockTableInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTableInserter) EXPECT() *MockTableInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTableInserter) Insert(key int64, row sql.Row) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", key, row)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTableInserterMockRecorder) Insert(key, row any) *MockTableInserterInsertCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTableInserter)(nil).Insert), key, row)
	return &MockTableInserterInsertCall{Call: call}
}

// MockTableInserterInsertCall wrap *gomock.Call
type MockTableInserterInsertCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTableInserterInsertCall) Return(arg0 error) *MockTableInserterInsertCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTableInserterInsertCall) Do(f func(int64, sql.Row) error) *MockTableInserterInsertCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTableInserterInsertCall) DoAndReturn(f func(int64, sql.Row) error) *MockTableInserterInsertCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
