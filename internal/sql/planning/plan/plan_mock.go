// Code generated by MockGen. DO NOT EDIT.
// Source: plan.go
//
// Generated by this command:
//
//	mockgen -typed -source=plan.go -destination ./plan_mock.go -package plan
//

// Package plan is a generated GoMock package.
package plan

import (
	reflect "reflect"

	sql "github.com/i-sevostyanov/NanoDB/internal/sql"
	gomock "go.uber.org/mock/gomock"
)

// MockNode is a mock of Node interface.
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode.
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance.
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// Columns mocks base method.
func (m *MockNode) Columns() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Columns indicates an expected call of Columns.
func (mr *MockNodeMockRecorder) Columns() *MockNodeColumnsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockNode)(nil).Columns))
	return &MockNodeColumnsCall{Call: call}
}

// MockNodeColumnsCall wrap *gomock.Call
type MockNodeColumnsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNodeColumnsCall) Return(arg0 []string) *MockNodeColumnsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNodeColumnsCall) Do(f func() []string) *MockNodeColumnsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNodeColumnsCall) DoAndReturn(f func() []string) *MockNodeColumnsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RowIter mocks base method.
func (m *MockNode) RowIter() (sql.RowIter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowIter")
	ret0, _ := ret[0].(sql.RowIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RowIter indicates an expected call of RowIter.
func (mr *MockNodeMockRecorder) RowIter() *MockNodeRowIterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowIter", reflect.TypeOf((*MockNode)(nil).RowIter))
	return &MockNodeRowIterCall{Call: call}
}

// MockNodeRowIterCall wrap *gomock.Call
type MockNodeRowIterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNodeRowIterCall) Return(arg0 sql.RowIter, arg1 error) *MockNodeRowIterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNodeRowIterCall) Do(f func() (sql.RowIter, error)) *MockNodeRowIterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNodeRowIterCall) DoAndReturn(f func() (sql.RowIter, error)) *MockNodeRowIterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
