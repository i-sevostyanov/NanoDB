package ast

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql/token"
)

type Node interface{}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Tree struct {
	Statements []Statement
}

type BadStatement struct {
	Type    token.Type
	Literal string
}

func (s *BadStatement) statementNode() {}

type SelectStatement struct {
	Result  []ResultStatement
	From    *FromStatement
	Where   *WhereStatement
	OrderBy *OrderByStatement
	Limit   *LimitStatement
	Offset  *OffsetStatement
}

func (s *SelectStatement) statementNode() {}

type ResultStatement struct {
	Alias Expression
	Expr  Expression
}

func (s *ResultStatement) statementNode() {}

type FromStatement struct {
	Table Expression
}

func (s *FromStatement) statementNode() {}

type WhereStatement struct {
	Expr Expression
}

func (s *WhereStatement) statementNode() {}

type OrderByStatement struct {
	Column Expression
	Order  Expression
}

func (s *OrderByStatement) statementNode() {}

type LimitStatement struct {
	Value Expression
}

func (s *LimitStatement) statementNode() {}

type OffsetStatement struct {
	Value Expression
}

func (s *OffsetStatement) statementNode() {}

type InsertStatement struct {
	Table   Expression
	Columns []IdentExpr
	Values  []Expression
}

func (s *InsertStatement) statementNode() {}

type UpdateStatement struct {
	Table Expression
	Set   []SetStatement
	Where *WhereStatement
}

func (s *UpdateStatement) statementNode() {}

type SetStatement struct {
	Column Expression
	Value  Expression
}

func (s *SetStatement) statementNode() {}

type DeleteStatement struct {
	Table Expression
	Where *WhereStatement
}

func (s *DeleteStatement) statementNode() {}

type CreateTableStatement struct {
	Table   Expression
	Columns []Column
}

func (s *CreateTableStatement) statementNode() {}

type Column struct {
	Name Expression
	Type token.Type
}

type DropTableStatement struct {
	Table Expression
}

func (s *DropTableStatement) statementNode() {}

type IdentExpr struct {
	Name string
}

func (e *IdentExpr) expressionNode() {}

type BinaryExpr struct {
	Left     Expression
	Operator token.Type
	Right    Expression
}

func (e *BinaryExpr) expressionNode() {}

type UnaryExpr struct {
	Operator token.Type
	Right    Expression
}

func (e *UnaryExpr) expressionNode() {}

type BadExpr struct {
	Type    token.Type
	Literal string
}

func (e *BadExpr) expressionNode() {}

type ScalarExpr struct {
	Type    token.Type
	Literal string
}

func (e *ScalarExpr) expressionNode() {}
