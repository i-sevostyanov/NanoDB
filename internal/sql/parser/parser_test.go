package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parser"
	"github.com/i-sevostyanov/NanoDB/internal/sql/token"
)

func TestParser_Select(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		ast   *ast.Tree
	}{
		{
			input: "SELECT id AS alias",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr:  &ast.IdentExpr{Name: "id"},
								Alias: &ast.IdentExpr{Name: "alias"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id AS alias, name",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr:  &ast.IdentExpr{Name: "id"},
								Alias: &ast.IdentExpr{Name: "alias"},
							},
							{
								Expr: &ast.IdentExpr{Name: "name"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id, salary, name AS alias",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{Name: "id"},
							},
							{
								Expr: &ast.IdentExpr{Name: "salary"},
							},
							{
								Expr:  &ast.IdentExpr{Name: "name"},
								Alias: &ast.IdentExpr{Name: "alias"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id AS PK, salary AS salary, name AS first_name",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr:  &ast.IdentExpr{Name: "id"},
								Alias: &ast.IdentExpr{Name: "PK"},
							},
							{
								Expr:  &ast.IdentExpr{Name: "salary"},
								Alias: &ast.IdentExpr{Name: "salary"},
							},
							{
								Expr:  &ast.IdentExpr{Name: "name"},
								Alias: &ast.IdentExpr{Name: "first_name"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id, 10*2 AS expr",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{Name: "id"},
							},
							{
								Expr: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
									Operator: token.Mul,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
								},
								Alias: &ast.IdentExpr{Name: "expr"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{Name: "id"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id, name",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{Name: "id"},
							},
							{
								Expr: &ast.IdentExpr{Name: "name"},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT 10+2*3",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
									Operator: token.Add,
									Right: &ast.BinaryExpr{
										Left: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "2",
										},
										Operator: token.Mul,
										Right: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "3",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT (10+2)*3",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.BinaryExpr{
									Left: &ast.BinaryExpr{
										Left: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "10",
										},
										Operator: token.Add,
										Right: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "2",
										},
									},
									Operator: token.Mul,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "3",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT 6+2^3*5-3+4/(10-2)%3",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.BinaryExpr{
									Left: &ast.BinaryExpr{
										Left: &ast.BinaryExpr{
											Left: &ast.ScalarExpr{
												Type:    token.Integer,
												Literal: "6",
											},
											Operator: token.Add,
											Right: &ast.BinaryExpr{
												Left: &ast.BinaryExpr{
													Left: &ast.ScalarExpr{
														Type:    token.Integer,
														Literal: "2",
													},
													Operator: token.Pow,
													Right: &ast.ScalarExpr{
														Type:    token.Integer,
														Literal: "3",
													},
												},
												Operator: token.Mul,
												Right: &ast.ScalarExpr{
													Type:    token.Integer,
													Literal: "5",
												},
											},
										},
										Operator: token.Sub,
										Right: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "3",
										},
									},
									Operator: token.Add,
									Right: &ast.BinaryExpr{
										Left: &ast.BinaryExpr{
											Left: &ast.ScalarExpr{
												Type:    token.Integer,
												Literal: "4",
											},
											Operator: token.Quo,
											Right: &ast.BinaryExpr{
												Left: &ast.ScalarExpr{
													Type:    token.Integer,
													Literal: "10",
												},
												Operator: token.Sub,
												Right: &ast.ScalarExpr{
													Type:    token.Integer,
													Literal: "2",
												},
											},
										},
										Operator: token.Rem,
										Right: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "3",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM table_name",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "table_name",
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad'",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "id"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
								},
								Operator: token.And,
								Right: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "name"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.String,
										Literal: "vlad",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "id"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
								},
								Operator: token.And,
								Right: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "name"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.String,
										Literal: "vlad",
									},
								},
							},
						},
						OrderBy: &ast.OrderByStatement{
							Column: &ast.IdentExpr{Name: "id"},
							Order:  &ast.IdentExpr{Name: "ASC"},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC LIMIT 99",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "id"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
								},
								Operator: token.And,
								Right: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "name"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.String,
										Literal: "vlad",
									},
								},
							},
						},
						OrderBy: &ast.OrderByStatement{
							Column: &ast.IdentExpr{Name: "id"},
							Order:  &ast.IdentExpr{Name: "ASC"},
						},
						Limit: &ast.LimitStatement{
							Value: &ast.ScalarExpr{
								Type:    token.Integer,
								Literal: "99",
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC LIMIT 99 OFFSET 10",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "id"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
								},
								Operator: token.And,
								Right: &ast.BinaryExpr{
									Left:     &ast.IdentExpr{Name: "name"},
									Operator: token.Equal,
									Right: &ast.ScalarExpr{
										Type:    token.String,
										Literal: "vlad",
									},
								},
							},
						},
						OrderBy: &ast.OrderByStatement{
							Column: &ast.IdentExpr{Name: "id"},
							Order:  &ast.IdentExpr{Name: "ASC"},
						},
						Limit: &ast.LimitStatement{
							Value: &ast.ScalarExpr{
								Type:    token.Integer,
								Literal: "99",
							},
						},
						Offset: &ast.OffsetStatement{
							Value: &ast.ScalarExpr{
								Type:    token.Integer,
								Literal: "10",
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT 2 ^ 3 ^ 4",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
									Operator: token.Pow,
									Right: &ast.BinaryExpr{
										Left: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "3",
										},
										Operator: token.Pow,
										Right: &ast.ScalarExpr{
											Type:    token.Integer,
											Literal: "4",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT -2, +2",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.UnaryExpr{
									Operator: token.Sub,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
								},
							},
							{
								Expr: &ast.UnaryExpr{
									Operator: token.Add,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers ORDER BY id",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.SelectStatement{
						Result: []ast.ResultStatement{
							{
								Expr: &ast.IdentExpr{
									Name: "id",
								},
							},
						},
						From: &ast.FromStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
						},
						OrderBy: &ast.OrderByStatement{
							Column: &ast.IdentExpr{Name: "id"},
							Order:  &ast.IdentExpr{Name: "ASC"},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			tree, errors := p.Parse()
			assert.Equal(t, test.ast, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})
	}

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name  string
			input string
			ast   *ast.Tree
		}{
			{
				name:  "unexpected statement",
				input: "SEL",
				ast: &ast.Tree{
					Statements: nil,
				},
			},
			{
				name:  "no columns specified",
				input: "SELECT",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{},
					},
				},
			},
			{
				name:  "no column alias specified",
				input: "SELECT id AS",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{},
					},
				},
			},
			{
				name:  "unexpected column alias",
				input: "SELECT id AS 7",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{},
					},
				},
			},
			{
				name:  "table name not specified",
				input: "SELECT id FROM",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
						},
					},
				},
			},
			{
				name:  "unexpected table name",
				input: "SELECT id FROM 9",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
						},
					},
				},
			},
			{
				name:  "empty 'where' expression",
				input: "SELECT id FROM customers WHERE",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "unfinished 'order by' statement",
				input: "SELECT id FROM customers WHERE id > 2 ORDER",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
							Where: &ast.WhereStatement{
								Expr: &ast.BinaryExpr{
									Left: &ast.IdentExpr{
										Name: "id",
									},
									Operator: token.GreaterThan,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
								},
							},
						},
					},
				},
			},
			{
				name:  "no column specified at 'order by' statement",
				input: "SELECT id FROM customers ORDER BY",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "unexpected 'order by' column type",
				input: "SELECT id FROM customers ORDER BY 9",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "'limit' value not specified",
				input: "SELECT id FROM customers LIMIT",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "unexpected 'limit' value type",
				input: "SELECT id FROM customers LIMIT abc",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "'offset' value not specified",
				input: "SELECT id FROM customers OFFSET",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "unexpected 'offset' value type",
				input: "SELECT id FROM customers OFFSET abc",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{
							Result: []ast.ResultStatement{
								{
									Expr: &ast.IdentExpr{
										Name: "id",
									},
								},
							},
							From: &ast.FromStatement{
								Table: &ast.IdentExpr{
									Name: "customers",
								},
							},
						},
					},
				},
			},
			{
				name:  "no close paren in group expr",
				input: "SELECT (10-2",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.SelectStatement{},
					},
				},
			},
		}

		for _, test := range tests {
			test := test

			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(test.input))
				tree, errors := p.Parse()

				assert.Equal(t, test.ast, tree)
				require.NotEmpty(t, errors)

				for _, err := range errors {
					assert.NotNil(t, err)
				}
			})
		}
	})
}

func TestParser_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		ast   *ast.Tree
	}{
		{
			input: "INSERT INTO customers (id, name, salary) VALUES (10, 'ivan', 10*2+1000)",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.InsertStatement{
						Table: &ast.IdentExpr{Name: "customers"},
						Columns: []ast.Expression{
							&ast.IdentExpr{Name: "id"},
							&ast.IdentExpr{Name: "name"},
							&ast.IdentExpr{Name: "salary"},
						},
						Values: []ast.Expression{
							&ast.ScalarExpr{
								Type:    token.Integer,
								Literal: "10",
							},
							&ast.ScalarExpr{
								Type:    token.String,
								Literal: "ivan",
							},
							&ast.BinaryExpr{
								Left: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
									Operator: token.Mul,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "2",
									},
								},
								Operator: token.Add,
								Right: &ast.ScalarExpr{
									Type:    token.Integer,
									Literal: "1000",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			tree, errors := p.Parse()
			assert.Equal(t, test.ast, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParser_Update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		ast   *ast.Tree
	}{
		{
			input: "UPDATE customers SET name = 'vlad', salary = 10*100 WHERE id = 1",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.UpdateStatement{
						Table: &ast.IdentExpr{
							Name: "customers",
						},
						Set: []ast.SetStatement{
							{
								Column: &ast.IdentExpr{
									Name: "name",
								},
								Value: &ast.ScalarExpr{
									Type:    token.String,
									Literal: "vlad",
								},
							},
							{
								Column: &ast.IdentExpr{
									Name: "salary",
								},
								Value: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
									Operator: token.Mul,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "100",
									},
								},
							},
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.IdentExpr{
									Name: "id",
								},
								Operator: token.Equal,
								Right: &ast.ScalarExpr{
									Type:    token.Integer,
									Literal: "1",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			tree, errors := p.Parse()
			assert.Equal(t, test.ast, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParser_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		ast   *ast.Tree
	}{
		{
			input: "DELETE FROM customers WHERE salary < 10*100",
			ast: &ast.Tree{
				Statements: []ast.Statement{
					&ast.DeleteStatement{
						Table: &ast.IdentExpr{
							Name: "customers",
						},
						Where: &ast.WhereStatement{
							Expr: &ast.BinaryExpr{
								Left: &ast.IdentExpr{
									Name: "salary",
								},
								Operator: token.LessThan,
								Right: &ast.BinaryExpr{
									Left: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "10",
									},
									Operator: token.Mul,
									Right: &ast.ScalarExpr{
										Type:    token.Integer,
										Literal: "100",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			tree, errors := p.Parse()
			assert.Equal(t, test.ast, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParser_Create(t *testing.T) {
	t.Parallel()

	t.Run("create database", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "CREATE DATABASE customers"
			expected := &ast.Tree{
				Statements: []ast.Statement{
					&ast.CreateDatabaseStatement{
						Name: &ast.IdentExpr{
							Name: "customers",
						},
					},
				},
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()
			assert.Equal(t, expected, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})

		t.Run("wrong query", func(t *testing.T) {
			t.Parallel()

			input := "CREATE DATABASE ^"
			expected := &ast.Tree{
				Statements: nil,
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()

			assert.Equal(t, expected, tree)
			assert.Len(t, errors, 1)
			assert.NotNil(t, errors[0])
		})
	})

	t.Run("create table", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			input string
			ast   *ast.Tree
		}{
			{
				input: "CREATE TABLE customers (id INTEGER, name STRING, salary FLOAT, is_active BOOLEAN)",
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.CreateTableStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
							Columns: []ast.Column{
								{
									Name: &ast.IdentExpr{
										Name: "id",
									},
									Type: token.Integer,
								},
								{
									Name: &ast.IdentExpr{
										Name: "name",
									},
									Type: token.String,
								},
								{
									Name: &ast.IdentExpr{
										Name: "salary",
									},
									Type: token.Float,
								},
								{
									Name: &ast.IdentExpr{
										Name: "is_active",
									},
									Type: token.Boolean,
								},
							},
						},
					},
				},
			},
			{
				input: `
				CREATE TABLE customers (
					id INTEGER PRIMARY KEY, 
					name STRING NULL, 
					salary FLOAT NOT NULL, 
					is_active BOOLEAN NOT NULL DEFAULT true,
				)
			`,
				ast: &ast.Tree{
					Statements: []ast.Statement{
						&ast.CreateTableStatement{
							Table: &ast.IdentExpr{
								Name: "customers",
							},
							Columns: []ast.Column{
								{
									Name: &ast.IdentExpr{
										Name: "id",
									},
									Type:       token.Integer,
									PrimaryKey: true,
								},
								{
									Name: &ast.IdentExpr{
										Name: "name",
									},
									Type: token.String,
								},
								{
									Name: &ast.IdentExpr{
										Name: "salary",
									},
									Type:    token.Float,
									NotNull: true,
								},
								{
									Name: &ast.IdentExpr{
										Name: "is_active",
									},
									Type:    token.Boolean,
									NotNull: true,
									Default: &ast.ScalarExpr{
										Type:    token.Boolean,
										Literal: "true",
									},
								},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			test := test

			t.Run(test.input, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(test.input))
				tree, errors := p.Parse()
				assert.Equal(t, test.ast, tree)

				for _, err := range errors {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("create unexpected", func(t *testing.T) {
		t.Parallel()

		input := "CREATE abc"
		expected := &ast.Tree{
			Statements: nil,
		}

		p := parser.New(lexer.New(input))
		tree, errors := p.Parse()

		assert.Equal(t, expected, tree)
		assert.Len(t, errors, 1)
		assert.NotNil(t, errors[0])
	})
}

func TestParser_Drop(t *testing.T) {
	t.Parallel()

	t.Run("drop database", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "DROP DATABASE customers"
			expected := &ast.Tree{
				Statements: []ast.Statement{
					&ast.DropDatabaseStatement{
						Name: &ast.IdentExpr{
							Name: "customers",
						},
					},
				},
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()
			assert.Equal(t, expected, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})

		t.Run("wrong database name", func(t *testing.T) {
			t.Parallel()

			input := "DROP DATABASE 9"
			expected := &ast.Tree{
				Statements: nil,
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()

			assert.Equal(t, expected, tree)
			assert.Len(t, errors, 1)
			assert.NotNil(t, errors[0])
		})
	})

	t.Run("drop table", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "DROP TABLE customers"
			expected := &ast.Tree{
				Statements: []ast.Statement{
					&ast.DropTableStatement{
						Table: &ast.IdentExpr{
							Name: "customers",
						},
					},
				},
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()
			assert.Equal(t, expected, tree)

			for _, err := range errors {
				assert.NoError(t, err)
			}
		})

		t.Run("wrong table name", func(t *testing.T) {
			t.Parallel()

			input := "DROP TABLE +"
			expected := &ast.Tree{
				Statements: nil,
			}

			p := parser.New(lexer.New(input))
			tree, errors := p.Parse()

			assert.Equal(t, expected, tree)
			assert.Len(t, errors, 1)
			assert.NotNil(t, errors[0])
		})
	})

	t.Run("drop unexpected", func(t *testing.T) {
		t.Parallel()

		input := "DROP abc"
		expected := &ast.Tree{
			Statements: nil,
		}

		p := parser.New(lexer.New(input))
		tree, errors := p.Parse()

		assert.Equal(t, expected, tree)
		assert.Len(t, errors, 1)
		assert.NotNil(t, errors[0])
	})
}
