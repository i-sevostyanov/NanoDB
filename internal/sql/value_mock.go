// Code generated by MockGen. DO NOT EDIT.
// Source: value.go

// Package sql is a generated GoMock package.
package sql

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValue is a mock of Value interface
type MockValue struct {
	ctrl     *gomock.Controller
	recorder *MockValueMockRecorder
}

// MockValueMockRecorder is the mock recorder for MockValue
type MockValueMockRecorder struct {
	mock *MockValue
}

// NewMockValue creates a new mock instance
func NewMockValue(ctrl *gomock.Controller) *MockValue {
	mock := &MockValue{ctrl: ctrl}
	mock.recorder = &MockValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValue) EXPECT() *MockValueMockRecorder {
	return m.recorder
}

// Raw mocks base method
func (m *MockValue) Raw() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Raw indicates an expected call of Raw
func (mr *MockValueMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockValue)(nil).Raw))
}

// Compare mocks base method
func (m *MockValue) Compare(x Value) (CompareType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", x)
	ret0, _ := ret[0].(CompareType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare
func (mr *MockValueMockRecorder) Compare(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockValue)(nil).Compare), x)
}

// UnaryPlus mocks base method
func (m *MockValue) UnaryPlus() (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnaryPlus")
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnaryPlus indicates an expected call of UnaryPlus
func (mr *MockValueMockRecorder) UnaryPlus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnaryPlus", reflect.TypeOf((*MockValue)(nil).UnaryPlus))
}

// UnaryMinus mocks base method
func (m *MockValue) UnaryMinus() (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnaryMinus")
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnaryMinus indicates an expected call of UnaryMinus
func (mr *MockValueMockRecorder) UnaryMinus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnaryMinus", reflect.TypeOf((*MockValue)(nil).UnaryMinus))
}

// Add mocks base method
func (m *MockValue) Add(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockValueMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockValue)(nil).Add), arg0)
}

// Sub mocks base method
func (m *MockValue) Sub(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sub", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sub indicates an expected call of Sub
func (mr *MockValueMockRecorder) Sub(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockValue)(nil).Sub), arg0)
}

// Mul mocks base method
func (m *MockValue) Mul(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mul", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mul indicates an expected call of Mul
func (mr *MockValueMockRecorder) Mul(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mul", reflect.TypeOf((*MockValue)(nil).Mul), arg0)
}

// Div mocks base method
func (m *MockValue) Div(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Div", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Div indicates an expected call of Div
func (mr *MockValueMockRecorder) Div(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Div", reflect.TypeOf((*MockValue)(nil).Div), arg0)
}

// Pow mocks base method
func (m *MockValue) Pow(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pow", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pow indicates an expected call of Pow
func (mr *MockValueMockRecorder) Pow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pow", reflect.TypeOf((*MockValue)(nil).Pow), arg0)
}

// Mod mocks base method
func (m *MockValue) Mod(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mod", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mod indicates an expected call of Mod
func (mr *MockValueMockRecorder) Mod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mod", reflect.TypeOf((*MockValue)(nil).Mod), arg0)
}

// Equal mocks base method
func (m *MockValue) Equal(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Equal indicates an expected call of Equal
func (mr *MockValueMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockValue)(nil).Equal), arg0)
}

// NotEqual mocks base method
func (m *MockValue) NotEqual(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotEqual", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotEqual indicates an expected call of NotEqual
func (mr *MockValueMockRecorder) NotEqual(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotEqual", reflect.TypeOf((*MockValue)(nil).NotEqual), arg0)
}

// GreaterThan mocks base method
func (m *MockValue) GreaterThan(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GreaterThan", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GreaterThan indicates an expected call of GreaterThan
func (mr *MockValueMockRecorder) GreaterThan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GreaterThan", reflect.TypeOf((*MockValue)(nil).GreaterThan), arg0)
}

// LessThan mocks base method
func (m *MockValue) LessThan(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LessThan", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LessThan indicates an expected call of LessThan
func (mr *MockValueMockRecorder) LessThan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LessThan", reflect.TypeOf((*MockValue)(nil).LessThan), arg0)
}

// GreaterOrEqual mocks base method
func (m *MockValue) GreaterOrEqual(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GreaterOrEqual", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GreaterOrEqual indicates an expected call of GreaterOrEqual
func (mr *MockValueMockRecorder) GreaterOrEqual(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GreaterOrEqual", reflect.TypeOf((*MockValue)(nil).GreaterOrEqual), arg0)
}

// LessOrEqual mocks base method
func (m *MockValue) LessOrEqual(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LessOrEqual", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LessOrEqual indicates an expected call of LessOrEqual
func (mr *MockValueMockRecorder) LessOrEqual(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LessOrEqual", reflect.TypeOf((*MockValue)(nil).LessOrEqual), arg0)
}

// And mocks base method
func (m *MockValue) And(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "And", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// And indicates an expected call of And
func (mr *MockValueMockRecorder) And(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "And", reflect.TypeOf((*MockValue)(nil).And), arg0)
}

// Or mocks base method
func (m *MockValue) Or(arg0 Value) (Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Or", arg0)
	ret0, _ := ret[0].(Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Or indicates an expected call of Or
func (mr *MockValueMockRecorder) Or(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*MockValue)(nil).Or), arg0)
}
