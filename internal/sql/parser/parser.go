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
	errors    []error
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
func (p *Parser) Parse() (*ast.Statements, []error) {
	var statements ast.Statements

	for p.token.Type != token.Semicolon && p.token.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			statements = append(statements, stmt)
		}

		p.nextToken()
	}

	return &statements, p.errors
}

func (p *Parser) nextToken() {
	p.token = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.token.Type {
	// DQL
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
		p.errorf("unexpected statement %q", p.token.Type)
		return nil
	}
}

func (p *Parser) parseCreateStatement() ast.Statement {
	p.nextToken()

	switch p.token.Type {
	case token.Database:
		p.nextToken()
		return p.parseCreateDatabaseStatement()
	case token.Table:
		p.nextToken()
		return p.parseCreateTableStatement()
	default:
		p.errorf("unexpected keyword at CREATE statement %q", p.token.Type)
		return nil
	}
}

func (p *Parser) parseDropStatement() ast.Statement {
	p.nextToken()

	switch p.token.Type {
	case token.Database:
		p.nextToken()
		return p.parseDropDatabaseStatement()
	case token.Table:
		p.nextToken()
		return p.parseDropTableStatement()
	default:
		p.errorf("unexpected keyword at DROP statement %q", p.token.Type)
		return nil
	}
}

func (p *Parser) parseSelectStatement() ast.Statement {
	p.nextToken()

	return &ast.SelectStatement{
		Result:  p.parseResultStatement(),
		From:    p.parseFromStatement(),
		Where:   p.parseWhereStatement(),
		OrderBy: p.parseOrderByStatement(),
		Limit:   p.parseLimitStatement(),
		Offset:  p.parseOffsetStatement(),
	}
}

func (p *Parser) parseInsertStatement() ast.Statement {
	p.nextToken()

	if !p.expect(token.Into) {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.InsertStatement{
		Table:   table,
		Columns: p.parseColumnsStatement(),
		Values:  p.parseValuesStatement(),
	}
}

func (p *Parser) parseUpdateStatement() ast.Statement {
	p.nextToken()

	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.UpdateStatement{
		Table: table,
		Set:   p.parseSetStatement(),
		Where: p.parseWhereStatement(),
	}
}

func (p *Parser) parseDeleteStatement() ast.Statement {
	p.nextToken()

	if !p.expect(token.From) {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.DeleteStatement{
		Table: table,
		Where: p.parseWhereStatement(),
	}
}

func (p *Parser) parseCreateDatabaseStatement() ast.Statement {
	if !p.expect(token.Ident) {
		return nil
	}

	database := p.parseIdent()
	p.nextToken()

	return &ast.CreateDatabaseStatement{
		Name: database,
	}
}

func (p *Parser) parseCreateTableStatement() ast.Statement {
	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.CreateTableStatement{
		Table:   table,
		Columns: p.parseColumns(),
	}
}

func (p *Parser) parseColumns() []ast.Column {
	if !p.expect(token.OpenParen) {
		return nil
	}

	columns := make([]ast.Column, 0)

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
		}

		if !p.expect(token.Ident) {
			return nil
		}

		column := ast.Column{}
		column.Name = p.parseIdent()

		p.nextToken()

		switch p.token.Type {
		case token.Integer, token.Float, token.String, token.Boolean:
			column.Type = p.token.Type
		default:
			p.errorf("unexpected column type: %s", p.token.Type)
			return nil
		}

		if p.peekToken.Type == token.Not {
			p.nextToken()
			p.nextToken()

			if !p.expect(token.Null) {
				return nil
			}

			column.NotNull = true
		}

		if p.peekToken.Type == token.Null {
			p.nextToken()
			p.nextToken()

			column.NotNull = false
		}

		if p.peekToken.Type == token.Default {
			p.nextToken()
			p.nextToken()
			column.Default = p.parseExprStatement()
		}

		if p.peekToken.Type == token.Primary {
			p.nextToken()
			p.nextToken()

			if !p.expect(token.Key) {
				return nil
			}

			column.PrimaryKey = true
		}

		columns = append(columns, column)

		p.nextToken()
	}

	if !p.expect(token.CloseParen) {
		return nil
	}

	p.nextToken()

	return columns
}

func (p *Parser) parseDropDatabaseStatement() ast.Statement {
	if !p.expect(token.Ident) {
		return nil
	}

	database := p.parseIdent()
	p.nextToken()

	return &ast.DropDatabaseStatement{
		Name: database,
	}
}

func (p *Parser) parseDropTableStatement() ast.Statement {
	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.DropTableStatement{
		Table: table,
	}
}

func (p *Parser) parseResultStatement() []ast.ResultStatement {
	var result []ast.ResultStatement

	for p.token.Type != token.EOF && p.token.Type != token.From {
		var (
			expr  ast.Expression
			alias ast.Expression
		)

		if expr = p.parseExprStatement(); expr == nil {
			p.errorf("expression in result must be not empty")
			return nil
		}

		if p.peekToken.Type == token.As {
			p.nextToken()
			p.nextToken()

			if !p.expect(token.Ident) {
				return nil
			}

			alias = p.parseIdent()
			p.nextToken()
		}

		result = append(result, ast.ResultStatement{
			Alias: alias,
			Expr:  expr,
		})

		p.nextToken()
	}

	if len(result) == 0 {
		p.errorf("no columns specified")
		return nil
	}

	return result
}

func (p *Parser) parseFromStatement() *ast.FromStatement {
	if p.token.Type != token.From {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Ident) {
		return nil
	}

	table := p.parseIdent()
	p.nextToken()

	return &ast.FromStatement{
		Table: table,
	}
}

func (p *Parser) parseWhereStatement() *ast.WhereStatement {
	if p.token.Type != token.Where {
		return nil
	}

	p.nextToken()
	expr := p.parseExprStatement()
	p.nextToken()

	if expr == nil {
		p.errorf("WHERE expression must not be empty")
		return nil
	}

	return &ast.WhereStatement{
		Expr: expr,
	}
}

func (p *Parser) parseOrderByStatement() *ast.OrderByStatement {
	if p.token.Type != token.Order {
		return nil
	}

	p.nextToken()

	if !p.expect(token.By) {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Ident) {
		return nil
	}

	var order ast.Expression

	column := p.parseIdent()

	switch p.peekToken.Type {
	case token.Asc, token.Desc:
		p.nextToken()
		order = p.parseIdent()
	default:
		order = &ast.IdentExpr{Name: token.Asc.String()}
	}

	p.nextToken()

	return &ast.OrderByStatement{
		Column: column,
		Order:  order,
	}
}

func (p *Parser) parseLimitStatement() *ast.LimitStatement {
	if p.token.Type != token.Limit {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Integer) {
		return nil
	}

	value := p.parseScalar()
	p.nextToken()

	return &ast.LimitStatement{
		Value: value,
	}
}

func (p *Parser) parseOffsetStatement() *ast.OffsetStatement {
	if p.token.Type != token.Offset {
		return nil
	}

	p.nextToken()

	if !p.expect(token.Integer) {
		return nil
	}

	value := p.parseScalar()
	p.nextToken()

	return &ast.OffsetStatement{
		Value: value,
	}
}

func (p *Parser) parseColumnsStatement() []ast.Expression {
	var columns []ast.Expression

	if !p.expect(token.OpenParen) {
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
		}

		if !p.expect(token.Ident) {
			return nil
		}

		column := p.parseIdent()
		columns = append(columns, column)

		p.nextToken()
	}

	if !p.expect(token.CloseParen) {
		return nil
	}

	p.nextToken()

	return columns
}

func (p *Parser) parseValuesStatement() []ast.Expression {
	var values []ast.Expression

	if !p.expect(token.Values) {
		return nil
	}

	p.nextToken()

	if !p.expect(token.OpenParen) {
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		expr := p.parseExprStatement()
		if expr == nil {
			p.errorf("expression must not be empty")
			return nil
		}

		values = append(values, expr)

		p.nextToken()
	}

	if !p.expect(token.CloseParen) {
		return nil
	}

	p.nextToken()

	return values
}

func (p *Parser) parseSetStatement() []ast.SetStatement {
	if !p.expect(token.Set) {
		return nil
	}

	fields := make([]ast.SetStatement, 0)

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.Where {
		if !p.expect(token.Ident) {
			return nil
		}

		column := p.parseIdent()
		p.nextToken()

		if !p.expect(token.Equal) {
			return nil
		}

		p.nextToken()

		value := p.parseExprStatement()
		if value == nil {
			p.errorf("expression must not be empty")
			return nil
		}

		fields = append(fields, ast.SetStatement{
			Column: column,
			Value:  value,
		})

		p.nextToken()
	}

	return fields
}

func (p *Parser) parseExprStatement() ast.Expression {
	expr := p.parseExpr(token.LowestPrecedence + 1)

	if p.peekToken.Type == token.Comma {
		p.nextToken()
	}

	return expr
}

func (p *Parser) parseExpr(precedence int) ast.Expression {
	var expr ast.Expression

	if expr = p.parseOperand(); expr == nil {
		return nil
	}

	for p.peekToken.Type != token.Comma && precedence < p.peekToken.Type.Precedence() {
		if !p.isInfixOperator(p.peekToken.Type) {
			return expr
		}

		p.nextToken()
		expr = p.parseBinaryExpr(expr)
	}

	return expr
}

func (p *Parser) parseOperand() ast.Expression {
	switch p.token.Type {
	case token.Ident:
		return p.parseIdent()
	case token.Integer, token.Float, token.String, token.Boolean:
		return p.parseScalar()
	case token.Add, token.Sub:
		return p.parseUnaryExpr()
	case token.OpenParen:
		return p.parseGroupExpr()
	default:
		p.errorf("unexpected operand %q", p.token.Type)
		return nil
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

func (p *Parser) parseIdent() ast.Expression {
	return &ast.IdentExpr{
		Name: p.token.Literal,
	}
}

func (p *Parser) parseScalar() ast.Expression {
	return &ast.ScalarExpr{
		Type:    p.token.Type,
		Literal: p.token.Literal,
	}
}

func (p *Parser) parseUnaryExpr() ast.Expression {
	operator := p.token.Type
	p.nextToken()
	right := p.parseExpr(operator.Precedence())

	return &ast.UnaryExpr{
		Operator: operator,
		Right:    right,
	}
}

func (p *Parser) parseBinaryExpr(left ast.Expression) ast.Expression {
	operator := p.token.Type
	precedence := operator.Precedence()

	if operator.IsRightAssociative() {
		precedence--
	}

	p.nextToken()
	right := p.parseExpr(precedence)

	return &ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (p *Parser) parseGroupExpr() ast.Expression {
	p.nextToken()
	expr := p.parseExpr(token.LowestPrecedence)
	p.nextToken()

	if !p.expect(token.CloseParen) {
		return nil
	}

	return expr
}

func (p *Parser) expect(tokenType token.Type) bool {
	if p.token.Type == tokenType {
		return true
	}

	p.errorf(
		"expected %q but found %q (%s) at column %d",
		tokenType.String(),
		p.token.Literal,
		p.token.Type,
		p.token.Offset,
	)
	p.nextToken()

	return false
}

func (p *Parser) errorf(format string, a ...interface{}) {
	p.errors = append(p.errors, fmt.Errorf(format, a...))
}
