package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type BinaryOp uint

const (
	// Comparison
	Equal BinaryOp = iota
	NotEqual
	LessThan
	GreaterThan
	LessThanOrEqual
	GreaterThanOrEqual
	// Logical
	And
	Or
	// Mathematical
	Add
	Sub
	Mul
	Div
	Mod
	Pow
)

type Binary struct {
	Operator BinaryOp
	Left     Node
	Right    Node
}

func (b Binary) Eval(row sql.Row) (sql.Value, error) {
	lvalue, err := b.Left.Eval(row)
	if err != nil {
		return nil, fmt.Errorf("binary: failed to eval left arg: %w", err)
	}

	rvalue, err := b.Right.Eval(row)
	if err != nil {
		return nil, fmt.Errorf("binary: failed to eval right arg: %w", err)
	}

	switch b.Operator {
	case Equal:
		return lvalue.Equal(rvalue)
	case NotEqual:
		return lvalue.NotEqual(rvalue)
	case LessThan:
		return lvalue.LessThan(rvalue)
	case GreaterThan:
		return lvalue.GreaterThan(rvalue)
	case LessThanOrEqual:
		return lvalue.LessOrEqual(rvalue)
	case GreaterThanOrEqual:
		return lvalue.GreaterOrEqual(rvalue)
	case And:
		return lvalue.And(rvalue)
	case Or:
		return lvalue.Or(rvalue)
	case Add:
		return lvalue.Add(rvalue)
	case Sub:
		return lvalue.Sub(rvalue)
	case Mul:
		return lvalue.Mul(rvalue)
	case Div:
		return lvalue.Div(rvalue)
	case Mod:
		return lvalue.Mod(rvalue)
	case Pow:
		return lvalue.Pow(rvalue)
	default:
		return nil, fmt.Errorf("unknown binary Operator: %q", b.Operator)
	}
}
