// Package ast declares the types used to represent syntax trees for the NanoDB's SQL dialect.
package ast

import (
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
)

// Node represents AST-node of the syntax tree for SQL query.
type Node any

// Statement represents syntax tree node of SQL statement (like: SELECT).
type Statement interface {
	Node
	statementNode()
}

// Expression represents syntax tree node of SQL expression (like: id < 10 AND id > 5).
type Expression interface {
	Node
	expressionNode()
}

// SelectStatement node represents a SELECT statement.
type SelectStatement struct {
	Result  []ResultStatement
	From    *FromStatement
	Where   *WhereStatement
	OrderBy *OrderByStatement
	Limit   *LimitStatement
	Offset  *OffsetStatement
}

// ResultStatement node represents a returning expression in a SELECT statement.
type ResultStatement struct {
	Alias string
	Expr  Expression
}

// FromStatement node represents a FROM statement.
type FromStatement struct {
	Table string
}

// WhereStatement node represents a WHERE statement.
type WhereStatement struct {
	Expr Expression
}

// OrderByStatement node represents an ORDER BY statement.
type OrderByStatement struct {
	Column    string
	Direction token.Type
}

// LimitStatement node represents a LIMIT statement.
type LimitStatement struct {
	Value Expression
}

// OffsetStatement node represents a OFFSET statement.
type OffsetStatement struct {
	Value Expression
}

// InsertStatement node represents a INSERT statement.
type InsertStatement struct {
	Table   string
	Columns []string
	Values  []Expression
}

// UpdateStatement node represents a UPDATE statement.
type UpdateStatement struct {
	Table string
	Set   []SetStatement
	Where *WhereStatement
}

// SetStatement node represents a key-value pair (column => value) in UPDATE statement.
type SetStatement struct {
	Column string
	Value  Expression
}

// DeleteStatement node represents a DELETE statement.
type DeleteStatement struct {
	Table string
	Where *WhereStatement
}

// CreateDatabaseStatement node represents a CREATE DATABASE statement.
type CreateDatabaseStatement struct {
	Database string
}

// DropDatabaseStatement node represents a DROP DATABASE statement.
type DropDatabaseStatement struct {
	Database string
}

// CreateTableStatement node represents a CREATE TABLE statement.
type CreateTableStatement struct {
	Table   string
	Columns []Column
}

// Column node represents a table column definition.
type Column struct {
	Name       string
	Type       token.Type
	Default    Expression
	Nullable   bool
	PrimaryKey bool
}

// DropTableStatement node represents a DROP TABLE statement.
type DropTableStatement struct {
	Table string
}

func (s *SelectStatement) statementNode()         {}
func (s *ResultStatement) statementNode()         {}
func (s *FromStatement) statementNode()           {}
func (s *WhereStatement) statementNode()          {}
func (s *OrderByStatement) statementNode()        {}
func (s *LimitStatement) statementNode()          {}
func (s *OffsetStatement) statementNode()         {}
func (s *InsertStatement) statementNode()         {}
func (s *UpdateStatement) statementNode()         {}
func (s *SetStatement) statementNode()            {}
func (s *DeleteStatement) statementNode()         {}
func (s *CreateDatabaseStatement) statementNode() {}
func (s *DropDatabaseStatement) statementNode()   {}
func (s *CreateTableStatement) statementNode()    {}
func (s *DropTableStatement) statementNode()      {}

// IdentExpr node represents an identifier.
type IdentExpr struct {
	Name string
}

// BinaryExpr node represents a binary expression.
type BinaryExpr struct {
	Left     Expression
	Operator token.Type
	Right    Expression
}

// UnaryExpr node represents a unary expression.
type UnaryExpr struct {
	Operator token.Type
	Right    Expression
}

// ScalarExpr node represents a literal of basic type.
type ScalarExpr struct {
	Type    token.Type
	Literal string
}

// AsteriskExpr node represents asterisk at `SELECT *` expression.
type AsteriskExpr struct{}

func (e *IdentExpr) expressionNode()    {}
func (e *BinaryExpr) expressionNode()   {}
func (e *UnaryExpr) expressionNode()    {}
func (e *ScalarExpr) expressionNode()   {}
func (e *AsteriskExpr) expressionNode() {}
