package parser

import (
	"fmt"

	"github.com/i-sevostyanov/NanoDB/internal/sql/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	token     token.Token
	peekToken token.Token
	errors    []error
}

func New(lx *lexer.Lexer) *Parser {
	return &Parser{
		lexer:     lx,
		token:     lx.NextToken(),
		peekToken: lx.NextToken(),
	}
}

func (p *Parser) Parse() (*ast.Tree, []error) {
	var statements []ast.Statement

	for p.token.Type != token.Semicolon && p.token.Type != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			statements = append(statements, stmt)
		}

		p.nextToken()
	}

	tree := &ast.Tree{
		Statements: statements,
	}

	return tree, p.errors
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
		return p.parseBadStatement()
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
		return p.parseBadStatement()
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
		return p.parseBadStatement()
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

	if p.token.Type != token.Into {
		p.expect(token.Into)
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.InsertStatement{
		Table:   &ast.IdentExpr{Name: table},
		Columns: p.parseColumnsStatement(),
		Values:  p.parseValuesStatement(),
	}
}

func (p *Parser) parseUpdateStatement() ast.Statement {
	p.nextToken()

	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.UpdateStatement{
		Table: &ast.IdentExpr{Name: table},
		Set:   p.parseSetStatement(),
		Where: p.parseWhereStatement(),
	}
}

func (p *Parser) parseDeleteStatement() ast.Statement {
	p.nextToken()

	if p.token.Type != token.From {
		p.expect(token.From)
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.DeleteStatement{
		Table: &ast.IdentExpr{Name: table},
		Where: p.parseWhereStatement(),
	}
}

func (p *Parser) parseCreateDatabaseStatement() ast.Statement {
	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.CreateDatabaseStatement{
		Table: &ast.IdentExpr{
			Name: table,
		},
	}
}

func (p *Parser) parseCreateTableStatement() ast.Statement {
	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.CreateTableStatement{
		Table:   &ast.IdentExpr{Name: table},
		Columns: p.parseColumns(),
	}
}

func (p *Parser) parseColumns() []ast.Column {
	columns := make([]ast.Column, 0)

	if p.token.Type != token.OpenParen {
		p.expect(token.OpenParen)
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
		}

		if p.token.Type != token.Ident {
			p.expect(token.Ident)
			return nil
		}

		column := p.token.Literal
		p.nextToken()

		switch p.token.Type {
		case token.Integer, token.Float, token.String, token.Boolean:
			columns = append(columns, ast.Column{
				Name: &ast.IdentExpr{Name: column},
				Type: p.token.Type,
			})

			p.nextToken()

			// FIXME: parse constraint
		default:
			// p.expect(token.Ident)
			return columns
		}
	}

	p.expect(token.CloseParen)

	return columns
}

func (p *Parser) parseDropDatabaseStatement() ast.Statement {
	var table ast.Expression

	switch p.token.Type {
	case token.Ident:
		table = p.parseIdent()
	default:
		// p.expect(token.Ident)
		table = p.parseBadExpr()
	}

	p.nextToken()

	return &ast.DropDatabaseStatement{
		Table: table,
	}
}

func (p *Parser) parseDropTableStatement() ast.Statement {
	var table ast.Expression

	switch p.token.Type {
	case token.Ident:
		table = p.parseIdent()
	default:
		// p.expect(token.Ident)
		table = p.parseBadExpr()
	}

	p.nextToken()

	return &ast.DropTableStatement{
		Table: table,
	}
}

func (p *Parser) parseResultStatement() []ast.ResultStatement {
	var result []ast.ResultStatement

	for p.token.Type != token.EOF && p.token.Type != token.From {
		if expr := p.parseExprStatement(); expr != nil {
			var alias ast.Expression

			if p.peekToken.Type == token.As {
				p.nextToken()
				p.nextToken()

				if p.token.Type == token.Ident {
					alias = &ast.IdentExpr{
						Name: p.token.Literal,
					}

					p.nextToken()
				}
			}

			result = append(result, ast.ResultStatement{
				Alias: alias,
				Expr:  expr,
			})
		}

		p.nextToken()
	}

	return result
}

func (p *Parser) parseFromStatement() *ast.FromStatement {
	if p.token.Type != token.From {
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Ident {
		return &ast.FromStatement{
			Table: &ast.BadExpr{
				Type:    p.token.Type,
				Literal: p.token.Literal,
			},
		}
	}

	table := p.token.Literal
	p.nextToken()

	return &ast.FromStatement{
		Table: &ast.IdentExpr{
			Name: table,
		},
	}
}

func (p *Parser) parseWhereStatement() *ast.WhereStatement {
	if p.token.Type != token.Where {
		return nil
	}

	p.nextToken()
	expr := p.parseExprStatement()
	p.nextToken()

	return &ast.WhereStatement{
		Expr: expr,
	}
}

func (p *Parser) parseOrderByStatement() *ast.OrderByStatement {
	if p.token.Type != token.Order {
		return nil
	}

	p.nextToken()

	if p.token.Type != token.By {
		p.expect(token.By)
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Ident {
		p.expect(token.Ident)
		return nil
	}

	column := p.token.Literal
	p.nextToken()

	// order (asc, desc)
	if p.token.Type != token.Asc && p.token.Type != token.Desc {
		p.expect(token.Asc)
		return nil
	}

	order := p.token.Literal
	p.nextToken()

	return &ast.OrderByStatement{
		Column: &ast.IdentExpr{Name: column},
		Order:  &ast.IdentExpr{Name: order},
	}
}

func (p *Parser) parseLimitStatement() *ast.LimitStatement {
	if p.token.Type != token.Limit {
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Integer {
		p.expect(token.Integer)
		return nil
	}

	tokenType := p.token.Type
	literal := p.token.Literal
	p.nextToken()

	return &ast.LimitStatement{
		Value: &ast.ScalarExpr{
			Type:    tokenType,
			Literal: literal,
		},
	}
}

func (p *Parser) parseOffsetStatement() *ast.OffsetStatement {
	if p.token.Type != token.Offset {
		return nil
	}

	p.nextToken()

	if p.token.Type != token.Integer {
		p.expect(token.Integer)
		return nil
	}

	tokenType := p.token.Type
	literal := p.token.Literal
	p.nextToken()

	return &ast.OffsetStatement{
		Value: &ast.ScalarExpr{
			Type:    tokenType,
			Literal: literal,
		},
	}
}

func (p *Parser) parseColumnsStatement() []ast.IdentExpr {
	var columns []ast.IdentExpr

	if p.token.Type != token.OpenParen {
		p.expect(token.OpenParen)
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if p.token.Type == token.Comma {
			p.nextToken()
			continue
		}

		if p.token.Type != token.Ident {
			p.expect(token.Ident)
			continue
		}

		columns = append(columns, ast.IdentExpr{
			Name: p.token.Literal,
		})

		p.nextToken()
	}

	p.expect(token.CloseParen)

	return columns
}

func (p *Parser) parseValuesStatement() []ast.Expression {
	var values []ast.Expression

	if p.token.Type != token.Values {
		p.expect(token.Values)
		return nil
	}

	p.nextToken()

	if p.token.Type != token.OpenParen {
		p.expect(token.OpenParen)
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.CloseParen {
		if expr := p.parseExprStatement(); expr != nil {
			values = append(values, expr)
		}

		p.nextToken()
	}

	p.expect(token.CloseParen)

	return values
}

func (p *Parser) parseSetStatement() []ast.SetStatement {
	fields := make([]ast.SetStatement, 0)

	if p.token.Type != token.Set {
		p.expect(token.Ident)
		return nil
	}

	p.nextToken()

	for p.token.Type != token.EOF && p.token.Type != token.Where {
		if p.token.Type != token.Ident {
			p.expect(token.Ident)
			break
		}

		column := p.token.Literal
		p.nextToken()

		if p.token.Type != token.Equal {
			p.expect(token.Equal)
			break
		}

		p.nextToken()

		expr := p.parseExprStatement()
		if expr == nil {
			break
		}

		fields = append(fields, ast.SetStatement{
			Column: &ast.IdentExpr{Name: column},
			Value:  expr,
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
		return p.parseBadExpr()
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

func (p *Parser) parseBadStatement() ast.Statement {
	return &ast.BadStatement{
		Type:    p.token.Type,
		Literal: p.token.Literal,
	}
}

func (p *Parser) parseBadExpr() ast.Expression {
	return &ast.BadExpr{
		Type:    p.token.Type,
		Literal: p.token.Literal,
	}
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
	p.nextToken()
	right := p.parseExpr(operator.Precedence())

	return &ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (p *Parser) parseGroupExpr() ast.Expression {
	p.nextToken()
	expr := p.parseExpr(token.LowestPrecedence)

	if p.peekToken.Type != token.CloseParen {
		p.expect(token.CloseParen)
		return nil
	}

	p.nextToken()

	return expr
}

func (p *Parser) expect(tokenType token.Type) {
	if p.token.Type != tokenType {
		err := fmt.Errorf("expected: %q, but found: %q", tokenType.String(), p.token.Type.String())
		p.errors = append(p.errors, err)
	}

	p.nextToken()
}
