package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/comparison"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/logical"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr/math"
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
		return nil, fmt.Errorf("binary: eval left arg: %w", err)
	}

	rvalue, err := b.Right.Eval(row)
	if err != nil {
		return nil, fmt.Errorf("binary: eval right arg: %w", err)
	}

	switch b.Operator {
	case Equal:
		return comparison.Equal(lvalue, rvalue)
	case NotEqual:
		return comparison.NotEqual(lvalue, rvalue)
	case LessThan:
		return comparison.LessThan(lvalue, rvalue)
	case GreaterThan:
		return comparison.GreaterThan(lvalue, rvalue)
	case LessThanOrEqual:
		return comparison.LessOrEqual(lvalue, rvalue)
	case GreaterThanOrEqual:
		return comparison.GreaterOrEqual(lvalue, rvalue)
	case And:
		return logical.And(lvalue, rvalue)
	case Or:
		return logical.Or(lvalue, rvalue)
	case Add:
		return math.Add(lvalue, rvalue)
	case Sub:
		return math.Sub(lvalue, rvalue)
	case Mul:
		return math.Mul(lvalue, rvalue)
	case Div:
		return math.Div(lvalue, rvalue)
	case Mod:
		return math.Mod(lvalue, rvalue)
	case Pow:
		return math.Pow(lvalue, rvalue)
	default:
		return nil, fmt.Errorf("unknown binary operator: %q", b.Operator)
	}
}
