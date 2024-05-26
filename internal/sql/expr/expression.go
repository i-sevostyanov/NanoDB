package expr

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
)

//go:generate go run go.uber.org/mock/mockgen -typed -source=expression.go -destination ./expression_mock.go -package expr

// Node is a combination of one or more SQL expressions.
type Node interface {
	String() string
	Eval(row sql.Row) (sql.Value, error)
}

func New(node ast.Expression, scheme sql.Scheme) (Node, error) {
	return walk(node, scheme)
}

func walk(node ast.Expression, scheme sql.Scheme) (Node, error) {
	switch expr := node.(type) {
	case *ast.IdentExpr:
		return columnExpr(expr, scheme)
	case *ast.BinaryExpr:
		return binaryExpr(expr, scheme)
	case *ast.UnaryExpr:
		return unaryExpr(expr, scheme)
	case *ast.ScalarExpr:
		return scalarExpr(expr)
	default:
		return nil, fmt.Errorf("unknown expression: %v", expr)
	}
}

func columnExpr(expr *ast.IdentExpr, scheme sql.Scheme) (Node, error) {
	if scheme == nil {
		return nil, fmt.Errorf("schema not provided")
	}

	definition, ok := scheme[expr.Name]
	if !ok {
		return nil, fmt.Errorf("column %q not exists", expr.Name)
	}

	column := Column{
		Name:     definition.Name,
		Position: definition.Position,
	}

	return column, nil
}

func binaryExpr(expr *ast.BinaryExpr, scheme sql.Scheme) (Node, error) {
	left, err := walk(expr.Left, scheme)
	if err != nil {
		return nil, fmt.Errorf("walk left arg of binary expr: %w", err)
	}

	right, err := walk(expr.Right, scheme)
	if err != nil {
		return nil, fmt.Errorf("walk right arg of binary expr: %w", err)
	}

	binary := &Binary{
		Operator: BinaryOp(expr.Operator.String()),
		Left:     left,
		Right:    right,
	}

	return binary, nil
}

func unaryExpr(expr *ast.UnaryExpr, scheme sql.Scheme) (Node, error) {
	var operator UnaryOp

	switch expr.Operator {
	case token.Add:
		operator = UnaryPlus
	case token.Sub:
		operator = UnaryMinus
	default:
		return nil, fmt.Errorf("unexpected unary operator: %s", expr.Operator)
	}

	operand, err := walk(expr.Right, scheme)
	if err != nil {
		return nil, fmt.Errorf("walk left arg of unary expr: %w", err)
	}

	node := &Unary{
		Operator: operator,
		Operand:  operand,
	}

	return node, nil
}

func scalarExpr(expr *ast.ScalarExpr) (Node, error) {
	switch expr.Type {
	case token.Integer:
		return NewInteger(expr.Literal)
	case token.Float:
		return NewFloat(expr.Literal)
	case token.Text:
		return NewString(expr.Literal)
	case token.Boolean:
		return NewBoolean(expr.Literal)
	case token.Null:
		return NewNull(), nil
	default:
		return nil, fmt.Errorf("unexpected scalar type: %s", expr.Type)
	}
}
