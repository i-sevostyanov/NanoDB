package sql

import (
	"fmt"
	"strconv"
)

type DataType uint

const (
	Null DataType = iota
	Integer
	Float
	Text
	Boolean
)

type CompareType int

const (
	Less    CompareType = -1
	Equal   CompareType = 0
	Greater CompareType = 1
)

func (t DataType) String() string {
	switch t {
	case Integer:
		return "integer"
	case Float:
		return "float"
	case Text:
		return "text"
	case Boolean:
		return "boolean"
	case Null:
		return "null"
	default:
		return fmt.Sprintf("DataType<%s>", strconv.Itoa(int(t)))
	}
}

//go:generate mockgen -source=value.go -destination ./value_mock.go -package sql

type Value interface {
	Raw() interface{}
	DataType() DataType
	Compare(x Value) (CompareType, error)
	UnaryPlus() (Value, error)
	UnaryMinus() (Value, error)
	Add(Value) (Value, error)
	Sub(Value) (Value, error)
	Mul(Value) (Value, error)
	Div(Value) (Value, error)
	Pow(Value) (Value, error)
	Mod(Value) (Value, error)
	Equal(Value) (Value, error)
	NotEqual(Value) (Value, error)
	GreaterThan(Value) (Value, error)
	LessThan(Value) (Value, error)
	GreaterOrEqual(Value) (Value, error)
	LessOrEqual(Value) (Value, error)
	And(Value) (Value, error)
	Or(Value) (Value, error)
}
