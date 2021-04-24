package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
)

type BinaryOp string

const (
	Equal              BinaryOp = "="
	NotEqual           BinaryOp = "!="
	LessThan           BinaryOp = "<"
	GreaterThan        BinaryOp = ">"
	LessThanOrEqual    BinaryOp = "<="
	GreaterThanOrEqual BinaryOp = ">="
	And                BinaryOp = "AND"
	Or                 BinaryOp = "OR"
	Add                BinaryOp = "+"
	Sub                BinaryOp = "-"
	Mul                BinaryOp = "*"
	Div                BinaryOp = "/"
	Mod                BinaryOp = "%"
	Pow                BinaryOp = "^"
)

func (o BinaryOp) String() string {
	return string(o)
}

type Binary struct {
	Operator BinaryOp
	Left     Node
	Right    Node
}

func (b Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator.String(), b.Right.String())
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
		return nil, fmt.Errorf("unknown binary operator: %q", b.Operator)
	}
}
