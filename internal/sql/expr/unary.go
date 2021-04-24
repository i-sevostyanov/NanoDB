package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type UnaryOp uint

const (
	UnaryPlus UnaryOp = iota
	UnaryMinus
)

func (o UnaryOp) String() string {
	switch o {
	case UnaryPlus:
		return ""
	default:
		return "-"
	}
}

type Unary struct {
	Operator UnaryOp
	Operand  Node
}

func (e *Unary) String() string {
	return fmt.Sprintf("(%s%s)", e.Operator.String(), e.Operand.String())
}

func (e *Unary) Eval(row sql.Row) (sql.Value, error) {
	value, err := e.Operand.Eval(row)
	if err != nil {
		return nil, fmt.Errorf("unary: failed to eval right arg: %w", err)
	}

	switch e.Operator {
	case UnaryPlus:
		return value.UnaryPlus()
	case UnaryMinus:
		return value.UnaryMinus()
	default:
		return nil, fmt.Errorf("unexpected unary operation: %v", e.Operator)
	}
}
