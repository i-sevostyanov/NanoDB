package expr

import (
	"fmt"
	"strconv"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
)

type Integer struct {
	value int64
}

func NewInteger(literal string) (Integer, error) {
	value, err := strconv.ParseInt(literal, 10, 64)
	if err != nil {
		return Integer{}, fmt.Errorf("failed to parse integer literal: %w", err)
	}

	node := Integer{
		value: value,
	}

	return node, nil
}

func (i Integer) Eval(_ sql.Row) (sql.Value, error) {
	return datatype.NewInteger(i.value), nil
}

type Float struct {
	value float64
}

func NewFloat(literal string) (Node, error) {
	value, err := strconv.ParseFloat(literal, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse float literal: %w", err)
	}

	node := Float{
		value: value,
	}

	return node, nil
}

func (f Float) Eval(_ sql.Row) (sql.Value, error) {
	return datatype.NewFloat(f.value), nil
}

type String struct {
	value string
}

func NewString(literal string) (Node, error) {
	node := String{
		value: literal,
	}

	return node, nil
}

func (s String) Eval(_ sql.Row) (sql.Value, error) {
	return datatype.NewString(s.value), nil
}

type Boolean struct {
	value bool
}

func NewBoolean(literal string) (Node, error) {
	value, err := strconv.ParseBool(literal)
	if err != nil {
		return nil, fmt.Errorf("failed to parse boolean literal: %w", err)
	}

	node := Boolean{
		value: value,
	}

	return node, nil
}

func (b Boolean) Eval(_ sql.Row) (sql.Value, error) {
	return datatype.NewBoolean(b.value), nil
}

type Null struct{}

func NewNull() Null {
	return Null{}
}

func (b Null) Eval(_ sql.Row) (sql.Value, error) {
	return datatype.NewNull(), nil
}
