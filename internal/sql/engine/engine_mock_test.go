// Code generated by MockGen. DO NOT EDIT.
// Source: engine.go

// Package engine_test is a generated GoMock package.
package engine_test

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

// MockParser is a mock of Parser interface
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockParser) Parse(sql string) (ast.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", sql)
	ret0, _ := ret[0].(ast.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockParserMockRecorder) Parse(sql interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), sql)
}

// MockPlanner is a mock of Planner interface
type MockPlanner struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerMockRecorder
}

// MockPlannerMockRecorder is the mock recorder for MockPlanner
type MockPlannerMockRecorder struct {
	mock *MockPlanner
}

// NewMockPlanner creates a new mock instance
func NewMockPlanner(ctrl *gomock.Controller) *MockPlanner {
	mock := &MockPlanner{ctrl: ctrl}
	mock.recorder = &MockPlannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlanner) EXPECT() *MockPlannerMockRecorder {
	return m.recorder
}

// Plan mocks base method
func (m *MockPlanner) Plan(database string, node ast.Node) (plan.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Plan", database, node)
	ret0, _ := ret[0].(plan.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Plan indicates an expected call of Plan
func (mr *MockPlannerMockRecorder) Plan(database, node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plan", reflect.TypeOf((*MockPlanner)(nil).Plan), database, node)
}
