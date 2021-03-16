// Package parser implements a parser for the NanoDB's SQL dialect.
package parser

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/token"
)

// Lexer converts a sequence of characters into a sequence of tokens.
type Lexer interface {
	NextToken() token.Token
}

// Parser takes a Lexer and builds an abstract syntax tree.
type Parser struct {
	lexer     Lexer
	token     token.Token
	peekToken token.Token
}

// New returns new Parser.
func New(lx Lexer) *Parser {
	return &Parser{
		lexer:     lx,
		token:     lx.NextToken(),
		peekToken: lx.NextToken(),
	}
}

// Parse parses the sql and returns a list of statements.
func (p *Parser) Parse() (*ast.Statements, error) {
	var statements ast.Statements

	for p.token.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)

		p.nextToken()
	}

	return &statements, nil
}

func (p *Parser) nextToken() {
	p.token = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.token.Type {
	// DML
	case token.Select:
		return p.parseSelectStatement()
	case token.Insert:
		return p.parseInsertStatement()
	case token.Update:
		return p.parseUpdateStatement()
	case token.Delete:
		return p.parseDeleteStatement()
	// DDL
	case token.Create:
		return p.parseCreateStatement()
	case token.Drop:
		return p.parseDropStatement()
	default:
		return nil, fmt.Errorf("unexpected statement %q", p.token.Type)
	}
}

func (p *Parser) parseCreateStatement() (ast.Statement, error) {
	p.nextToken()

	switch p.token.Type {
	case token.Database:
		p.nextToken()
		return p.parseCreateDatabaseStatement()
	case token.Table:
		p.nextToken()
		return p.parseCreateTableStatement()
	default:
		return nil, fmt.Errorf("unexpected keyword after CREATE statement %q", p.token.Type)
	}
}

func (p *Parser) parseDropStatement() (ast.Statement, error) {
	p.nextToken()

	switch p.token.Type {
	case token.Database:
		p.nextToken()
		return p.parseDropDatabaseStatement()
	case token.Table:
		p.nextToken()
		return p.parseDropTableStatement()
	default:
		return nil, fmt.Errorf("unexpected keyword after DROP statement %q", p.token.Type)
	}
}

func (p *Parser) parseSelectStatement() (ast.Statement, error) {
	p.nextToken()

	result, err := p.parseResultStatement()
	if err != nil {
		return nil, err
	}

	from, err := p.parseFromStatement()
	if err != nil {
		return nil, err
	}

	where, err := p.parseWhereStatement()
	if err != nil {
		return nil, err
	}

	order, err := p.parseOrderByStatement()
	if err != nil {
		return nil, err
	}

	limit, err := p.parseLimitStatement()
	if err != nil {
		return nil, err
	}

	offset, err := p.parseOffsetStatement()
	if err != nil {
		return nil, err
	}

	selectStmt := ast.SelectStatement{
		Result:  result,
		From:    from,
		Where:   where,
		OrderBy: order,
		Limit:   limit,
		Offset:  offset,
	}

	return &selectStmt, nil
}

func (p *Parser) parseInsertStatement() (ast.Statement, error) {
	p.nextToken()

	if err := p.expect(token.Into); err != nil {
		return nil, err
	}

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	columns, err := p.parseColumnsStatement()
	if err != nil {
		return nil, err
	}

	values, err := p.parseValuesStatement()
	if err != nil {
		return nil, err
	}

	insert := ast.InsertStatement{
		Table:   table,
		Columns: columns,
		Values:  values,
	}

	return &insert, nil
}

func (p *Parser) parseUpdateStatement() (ast.Statement, error) {
	p.nextToken()

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	set, err := p.parseSetStatement()
	if err != nil {
		return nil, err
	}

	where, err := p.parseWhereStatement()
	if err != nil {
		return nil, err
	}

	update := ast.UpdateStatement{
		Table: table,
		Set:   set,
		Where: where,
	}

	return &update, nil
}

func (p *Parser) parseDeleteStatement() (ast.Statement, error) {
	p.nextToken()

	if err := p.expect(token.From); err != nil {
		return nil, err
	}

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	where, err := p.parseWhereStatement()
	if err != nil {
		return nil, err
	}

	deleteStmt := ast.DeleteStatement{
		Table: table,
		Where: where,
	}

	return &deleteStmt, nil
}

func (p *Parser) parseCreateDatabaseStatement() (ast.Statement, error) {
	database, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	create := ast.CreateDatabaseStatement{
		Name: database,
	}

	return &create, nil
}

func (p *Parser) parseCreateTableStatement() (ast.Statement, error) {
	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	columns, err := p.parseColumnsDefinition()
	if err != nil {
		return nil, err
	}

	create := ast.CreateTableStatement{
		Table:   table,
		Columns: columns,
	}

	return &create, nil
}

func (p *Parser) parseColumnsDefinition() ([]ast.Column, error) {
	if err := p.expect(token.OpenParen); err != nil {
		return nil, err
	}

	columns := make([]ast.Column, 0)

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
		}

		column, err := p.parseColumnDefinition()
		if err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}

	if err := p.expect(token.CloseParen); err != nil {
		return nil, err
	}

	return columns, nil
}

func (p *Parser) parseColumnDefinition() (ast.Column, error) {
	name, err := p.parseIdent()
	if err != nil {
		return ast.Column{}, err
	}

	columnType, err := p.parseColumnType()
	if err != nil {
		return ast.Column{}, err
	}

	notNull, err := p.parseColumnNullable()
	if err != nil {
		return ast.Column{}, err
	}

	defaultExpr, err := p.parseColumnDefault()
	if err != nil {
		return ast.Column{}, err
	}

	pk, err := p.parseColumnPrimaryKey()
	if err != nil {
		return ast.Column{}, err
	}

	column := ast.Column{
		Name:       name,
		Type:       columnType,
		Default:    defaultExpr,
		NotNull:    notNull,
		PrimaryKey: pk,
	}

	return column, nil
}

func (p *Parser) parseColumnType() (token.Type, error) {
	switch p.token.Type {
	case token.Integer, token.Float, token.String, token.Boolean:
		columnType := p.token.Type
		p.nextToken()

		return columnType, nil
	}

	return token.Illegal, fmt.Errorf("unexpected column type: %q", p.token.Type)
}

func (p *Parser) parseColumnNullable() (bool, error) {
	if p.token.Type == token.Null {
		p.nextToken()
		return false, nil
	}

	if p.token.Type == token.Not {
		p.nextToken()

		if err := p.expect(token.Null); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (p *Parser) parseColumnDefault() (ast.Expression, error) {
	if p.token.Type != token.Default {
		return nil, nil
	}

	p.nextToken()

	expr, err := p.parseExprStatement()
	if err != nil {
		return nil, err
	}

	p.nextToken()

	return expr, nil
}

func (p *Parser) parseColumnPrimaryKey() (bool, error) {
	if p.token.Type != token.Primary {
		return false, nil
	}

	p.nextToken()

	if err := p.expect(token.Key); err != nil {
		return false, err
	}

	p.nextToken()

	return true, nil
}

func (p *Parser) parseDropDatabaseStatement() (ast.Statement, error) {
	database, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	drop := ast.DropDatabaseStatement{
		Name: database,
	}

	return &drop, nil
}

func (p *Parser) parseDropTableStatement() (ast.Statement, error) {
	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	drop := ast.DropTableStatement{
		Table: table,
	}

	return &drop, nil
}

func (p *Parser) parseResultStatement() ([]ast.ResultStatement, error) {
	var results []ast.ResultStatement

	for p.token.Type != token.EOF && p.token.Type != token.From {
		result, err := p.parseResult()
		if err != nil {
			return nil, err
		}

		results = append(results, result)

		p.nextToken()
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no columns specified")
	}

	return results, nil
}

func (p *Parser) parseResult() (ast.ResultStatement, error) {
	var (
		result ast.ResultStatement
		err    error
	)

	result.Expr, err = p.parseExprStatement()
	if err != nil {
		return ast.ResultStatement{}, err
	}

	if result.Expr == nil {
		return ast.ResultStatement{}, fmt.Errorf("expression in result must be not empty")
	}

	if p.peekToken.Type != token.As {
		return result, nil
	}

	p.nextToken()
	p.nextToken()

	result.Alias, err = p.parseIdent()
	if err != nil {
		return ast.ResultStatement{}, err
	}

	return result, nil
}

func (p *Parser) parseFromStatement() (*ast.FromStatement, error) {
	if p.token.Type != token.From {
		return nil, nil
	}

	p.nextToken()

	table, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	from := ast.FromStatement{
		Table: table,
	}

	return &from, nil
}

func (p *Parser) parseWhereStatement() (*ast.WhereStatement, error) {
	if p.token.Type != token.Where {
		return nil, nil
	}

	p.nextToken()

	expr, err := p.parseExprStatement()
	if err != nil {
		return nil, err
	}

	if expr == nil {
		return nil, fmt.Errorf("WHERE expression must not be empty")
	}

	p.nextToken()

	where := ast.WhereStatement{
		Expr: expr,
	}

	return &where, nil
}

func (p *Parser) parseOrderByStatement() (*ast.OrderByStatement, error) {
	if p.token.Type != token.Order {
		return nil, nil
	}

	p.nextToken()

	if err := p.expect(token.By); err != nil {
		return nil, err
	}

	column, err := p.parseIdent()
	if err != nil {
		return nil, err
	}

	direction := &ast.IdentExpr{
		Name: token.Asc.String(),
	}

	switch p.token.Type {
	case token.Asc, token.Desc:
		direction.Name = p.token.Type.String()
		p.nextToken()
	}

	order := ast.OrderByStatement{
		Column:    column,
		Direction: direction,
	}

	return &order, nil
}

func (p *Parser) parseLimitStatement() (*ast.LimitStatement, error) {
	if p.token.Type != token.Limit {
		return nil, nil
	}

	p.nextToken()

	value, err := p.parseScalar(token.Integer)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	limit := ast.LimitStatement{
		Value: value,
	}

	return &limit, nil
}

func (p *Parser) parseOffsetStatement() (*ast.OffsetStatement, error) {
	if p.token.Type != token.Offset {
		return nil, nil
	}

	p.nextToken()

	value, err := p.parseScalar(token.Integer)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	offset := ast.OffsetStatement{
		Value: value,
	}

	return &offset, nil
}

func (p *Parser) parseColumnsStatement() ([]ast.Expression, error) {
	var columns []ast.Expression

	if err := p.expect(token.OpenParen); err != nil {
		return nil, err
	}

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
		}

		column, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}

	if err := p.expect(token.CloseParen); err != nil {
		return nil, err
	}

	return columns, nil
}

func (p *Parser) parseValuesStatement() ([]ast.Expression, error) {
	var values []ast.Expression

	if err := p.expect(token.Values); err != nil {
		return nil, err
	}

	if err := p.expect(token.OpenParen); err != nil {
		return nil, err
	}

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		expr, err := p.parseExprStatement()
		if err != nil {
			return nil, err
		}

		if expr == nil {
			return nil, fmt.Errorf("expression must not be empty")
		}

		values = append(values, expr)

		p.nextToken()
	}

	if err := p.expect(token.CloseParen); err != nil {
		return nil, err
	}

	return values, nil
}

func (p *Parser) parseSetStatement() ([]ast.SetStatement, error) {
	if err := p.expect(token.Set); err != nil {
		return nil, err
	}

	fields := make([]ast.SetStatement, 0)

	for p.token.Type != token.EOF && p.token.Type != token.Where {
		column, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		if err = p.expect(token.Equal); err != nil {
			return nil, err
		}

		value, err := p.parseExprStatement()
		if err != nil {
			return nil, err
		}

		if value == nil {
			return nil, fmt.Errorf("expression must not be empty")
		}

		fields = append(fields, ast.SetStatement{
			Column: column,
			Value:  value,
		})

		p.nextToken()
	}

	return fields, nil
}

func (p *Parser) parseExprStatement() (ast.Expression, error) {
	expr, err := p.parseExpr(token.LowestPrecedence + 1)
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type == token.Comma {
		p.nextToken()
	}

	return expr, nil
}

func (p *Parser) parseExpr(precedence int) (ast.Expression, error) {
	expr, err := p.parseOperand()
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type != token.Comma && precedence < p.peekToken.Type.Precedence() {
		if !p.isInfixOperator(p.peekToken.Type) {
			return expr, nil
		}

		p.nextToken()

		expr, err = p.parseBinaryExpr(expr)
		if err != nil {
			return nil, err
		}
	}

	return expr, nil
}

func (p *Parser) parseOperand() (ast.Expression, error) {
	switch p.token.Type {
	case token.Ident:
		return &ast.IdentExpr{Name: p.token.Literal}, nil
	case token.Integer, token.Float, token.String, token.Boolean:
		return p.parseScalar(p.token.Type)
	case token.Add, token.Sub:
		return p.parseUnaryExpr()
	case token.OpenParen:
		return p.parseGroupExpr()
	default:
		return nil, fmt.Errorf("unexpected operand %q", p.token.Type)
	}
}

func (p *Parser) isInfixOperator(t token.Type) bool {
	switch t {
	case token.Add,
		token.Sub,
		token.Mul,
		token.Quo,
		token.Rem,
		token.Pow,
		token.Equal,
		token.LessThan,
		token.GreaterThan,
		token.NotEqual,
		token.LessThanOrEqual,
		token.GreaterThanOrEqual,
		token.And,
		token.Or,
		token.Not:
		return true
	}

	return false
}

func (p *Parser) parseIdent() (ast.Expression, error) {
	if p.token.Type != token.Ident {
		return nil, fmt.Errorf("unexpected token %q", p.token.Type)
	}

	ident := ast.IdentExpr{
		Name: p.token.Literal,
	}

	p.nextToken()

	return &ident, nil
}

func (p *Parser) parseScalar(expected token.Type) (ast.Expression, error) {
	if p.token.Type != expected {
		return nil, fmt.Errorf("unexpected scalar type %q", p.token.Type)
	}

	scalar := ast.ScalarExpr{
		Type:    p.token.Type,
		Literal: p.token.Literal,
	}

	return &scalar, nil
}

func (p *Parser) parseUnaryExpr() (ast.Expression, error) {
	operator := p.token.Type
	p.nextToken()

	right, err := p.parseExpr(operator.Precedence())
	if err != nil {
		return nil, err
	}

	expr := ast.UnaryExpr{
		Operator: operator,
		Right:    right,
	}

	return &expr, nil
}

func (p *Parser) parseBinaryExpr(left ast.Expression) (ast.Expression, error) {
	operator := p.token.Type
	precedence := operator.Precedence()

	if operator.IsRightAssociative() {
		precedence--
	}

	p.nextToken()

	right, err := p.parseExpr(precedence)
	if err != nil {
		return nil, err
	}

	expr := ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}

	return &expr, nil
}

func (p *Parser) parseGroupExpr() (ast.Expression, error) {
	p.nextToken()

	expr, err := p.parseExpr(token.LowestPrecedence)
	if err != nil {
		return nil, err
	}

	p.nextToken()

	if p.token.Type != token.CloseParen {
		return nil, fmt.Errorf("expected %q but found %q", token.CloseParen, p.token.Type)
	}

	return expr, nil
}

func (p *Parser) expect(tokenType token.Type) error {
	defer p.nextToken()

	if p.token.Type == tokenType {
		return nil
	}

	return fmt.Errorf(
		"expected %q but found %q (%s) at column %d",
		tokenType.String(),
		p.token.Literal,
		p.token.Type,
		p.token.Offset,
	)
}
