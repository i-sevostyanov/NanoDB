package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/lexer"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/parser"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
)

func TestParser_Select(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "SELECT id AS alias",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr:  &ast.IdentExpr{Name: "id"},
						Alias: &ast.IdentExpr{Name: "alias"},
					},
				},
			},
		},
		{
			input: "SELECT *, id FROM users",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.AsteriskExpr{},
					},
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "users",
				},
			},
		},
		{
			input: "SELECT id AS alias, name",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT id, salary, name AS alias",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT id AS PK, salary AS salary, name AS first_name",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT id, 10*2 AS expr",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT id",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{Name: "id"},
					},
				},
			},
		},
		{
			input: "SELECT id, name",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT 10+2*3",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT (10+2)*3",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT 6+2^3*5-3+4/(10-2)%3",
			stmt: &ast.SelectStatement{
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
									Operator: token.Div,
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
								Operator: token.Mod,
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
		{
			input: "SELECT id FROM table_name",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "table_name",
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad'",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
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
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
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
					Column:    "id",
					Direction: token.Asc,
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC LIMIT 99",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
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
					Column:    "id",
					Direction: token.Asc,
				},
				Limit: &ast.LimitStatement{
					Value: &ast.ScalarExpr{
						Type:    token.Integer,
						Literal: "99",
					},
				},
			},
		},
		{
			input: "SELECT id FROM customers WHERE id = 10 AND name = 'vlad' ORDER BY id ASC LIMIT 99 OFFSET 10",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
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
					Column:    "id",
					Direction: token.Asc,
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
		{
			input: "SELECT 2 ^ 3 ^ 4",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT -2, +2",
			stmt: &ast.SelectStatement{
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
		{
			input: "SELECT id FROM customers ORDER BY id",
			stmt: &ast.SelectStatement{
				Result: []ast.ResultStatement{
					{
						Expr: &ast.IdentExpr{
							Name: "id",
						},
					},
				},
				From: &ast.FromStatement{
					Table: "customers",
				},
				OrderBy: &ast.OrderByStatement{
					Column:    "id",
					Direction: token.Asc,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			stmts, err := p.Parse()
			require.NoError(t, err)
			assert.Equal(t, test.stmt, stmts)
		})
	}

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()

		inputs := []string{
			"SEL",
			"SELECT",
			"SELECT id AS",
			"SELECT id AS 7",
			"SELECT id FROM",
			"SELECT id FROM 9",
			"SELECT id FROM customers WHERE",
			"SELECT id FROM customers WHERE id > 2 ORDER",
			"SELECT id FROM customers ORDER BY",
			"SELECT id FROM customers ORDER BY 9",
			"SELECT id FROM customers LIMIT",
			"SELECT id FROM customers LIMIT abc",
			"SELECT id FROM customers OFFSET",
			"SELECT id FROM customers OFFSET abc",
			"SELECT (10-2",
		}

		for _, input := range inputs {
			input := input

			t.Run(input, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(input))
				stmts, err := p.Parse()

				require.NotNil(t, err)
				assert.Nil(t, stmts)
			})
		}
	})
}

func TestParser_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "INSERT INTO customers (id, name, salary) VALUES (10, 'ivan', 10*2+1000)",
			stmt: &ast.InsertStatement{
				Table: "customers",
				Columns: []string{
					"id",
					"name",
					"salary",
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
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			stmts, err := p.Parse()
			require.NoError(t, err)
			assert.Equal(t, test.stmt, stmts)
		})
	}

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()

		inputs := []string{
			"INSERT",
			"INSERT INTO",
			"INSERT INTO +",
			"INSERT INTO customers",
			"INSERT INTO customers id",
			"INSERT INTO customers (",
			"INSERT INTO customers (id,",
			"INSERT INTO customers (id,)",
			"INSERT INTO customers (id, name)",
			"INSERT INTO customers (id, name) VALUES",
			"INSERT INTO customers (id, name) VALUES id",
			"INSERT INTO customers (id, name) VALUES 1, 2",
			"INSERT INTO customers (id, name) VALUES (1, ",
			"INSERT INTO customers (id, name) VALUES (1, 'UA502',",
			"INSERT INTO customers (id, name) VALUES (+)",
		}

		for _, input := range inputs {
			input := input

			t.Run(input, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(input))
				stmts, err := p.Parse()

				require.NotNil(t, err)
				assert.Nil(t, stmts)
			})
		}
	})
}

func TestParser_Update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "UPDATE customers SET name = 'vlad', salary = 10*100 WHERE id = 1",
			stmt: &ast.UpdateStatement{
				Table: "customers",
				Set: []ast.SetStatement{
					{
						Column: "name",
						Value: &ast.ScalarExpr{
							Type:    token.String,
							Literal: "vlad",
						},
					},
					{
						Column: "salary",
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
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			stmts, err := p.Parse()
			require.NoError(t, err)
			assert.Equal(t, test.stmt, stmts)
		})
	}

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()

		inputs := []string{
			"UPDATE",
			"UPDATE +",
			"UPDATE customers",
			"UPDATE customers 9",
			"UPDATE customers WHERE id = 1",
			"UPDATE customers name",
			"UPDATE customers SET 9",
			"UPDATE customers SET name + ",
			"UPDATE customers SET name = ",
			"UPDATE customers SET name = +",
			"UPDATE customers SET name = 'max' id = 1",
			"UPDATE customers SET name = 'max' WHERE",
			"UPDATE customers SET name = 'max' WHERE id =)",
			"UPDATE customers SET name = 'max' WHERE +",
		}

		for _, input := range inputs {
			input := input

			t.Run(input, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(input))
				stmts, err := p.Parse()

				require.NotNil(t, err)
				assert.Nil(t, stmts)
			})
		}
	})
}

func TestParser_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		stmt  ast.Statement
	}{
		{
			input: "DELETE FROM customers WHERE salary < 10*100",
			stmt: &ast.DeleteStatement{
				Table: "customers",
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
	}

	for _, test := range tests {
		test := test

		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			p := parser.New(lexer.New(test.input))
			stmts, err := p.Parse()
			require.NoError(t, err)
			assert.Equal(t, test.stmt, stmts)
		})
	}

	t.Run("returns error", func(t *testing.T) {
		t.Parallel()

		inputs := []string{
			"DELETE",
			"DELETE FROM",
			"DELETE FROM +",
			"DELETE FROM customers WHERE SELECT",
			"DELETE FROM customers WHERE",
			"DELETE FROM customers WHERE (",
		}

		for _, input := range inputs {
			input := input

			t.Run(input, func(t *testing.T) {
				t.Parallel()

				p := parser.New(lexer.New(input))
				stmts, err := p.Parse()

				require.NotNil(t, err)
				assert.Nil(t, stmts)
			})
		}
	})
}

func TestParser_Create(t *testing.T) {
	t.Parallel()

	t.Run("create database", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "CREATE DATABASE customers"
			expected := &ast.CreateDatabaseStatement{
				Database: "customers",
			}

			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			require.NoError(t, err)
			assert.Equal(t, expected, stmts)
		})

		t.Run("wrong query", func(t *testing.T) {
			t.Parallel()

			input := "CREATE DATABASE ^"
			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			assert.Nil(t, stmts)
			assert.NotNil(t, err)
		})
	})

	t.Run("create table", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			input string
			stmt  ast.Statement
		}{
			{
				input: "CREATE TABLE customers (id INTEGER, name STRING, salary FLOAT, is_active BOOLEAN)",
				stmt: &ast.CreateTableStatement{
					Table: "customers",
					Columns: []ast.Column{
						{
							Name: "id",
							Type: token.Integer,
						},
						{
							Name: "name",
							Type: token.String,
						},
						{
							Name: "salary",
							Type: token.Float,
						},
						{
							Name: "is_active",
							Type: token.Boolean,
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
				stmt: &ast.CreateTableStatement{
					Table: "customers",
					Columns: []ast.Column{
						{
							Name:       "id",
							Type:       token.Integer,
							PrimaryKey: true,
							Nullable:   false,
						},
						{
							Name:     "name",
							Type:     token.String,
							Nullable: true,
						},
						{
							Name:     "salary",
							Type:     token.Float,
							Nullable: false,
						},
						{
							Name:     "is_active",
							Type:     token.Boolean,
							Nullable: false,
							Default: &ast.ScalarExpr{
								Type:    token.Boolean,
								Literal: "true",
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
				stmts, err := p.Parse()

				require.NoError(t, err)
				assert.Equal(t, test.stmt, stmts)
			})
		}

		t.Run("returns error", func(t *testing.T) {
			t.Parallel()

			inputs := []string{
				"CREATE",
				"CREATE TABLE",
				"CREATE TABLE 666",
				"CREATE TABLE customers id INTEGER",
				"CREATE TABLE customers (99 INTEGER)",
				"CREATE TABLE customers (id",
				"CREATE TABLE customers (id INT)",
				"CREATE TABLE customers (id INT)",
				"CREATE TABLE customers (id INTEGER",
				"CREATE TABLE customers (id INTEGER PRIMARY)",
				"CREATE TABLE customers (id INTEGER PRIMARY NULL)",
				"CREATE TABLE customers (id INTEGER DEFAULT &)",
				"CREATE TABLE customers (id INTEGER NOT)",
				"CREATE TABLE customers (id INTEGER DEFAULT NOT)",
				"CREATE TABLE customers (id INTEGER DEFAULT NOT KEY)",
			}

			for _, input := range inputs {
				input := input

				t.Run(input, func(t *testing.T) {
					t.Parallel()

					p := parser.New(lexer.New(input))
					stmts, err := p.Parse()

					require.NotNil(t, err)
					assert.Nil(t, stmts)
				})
			}
		})
	})

	t.Run("create unexpected", func(t *testing.T) {
		t.Parallel()

		input := "CREATE abc"
		p := parser.New(lexer.New(input))
		stmts, err := p.Parse()

		require.NotNil(t, err)
		assert.Nil(t, stmts)
	})
}

func TestParser_Drop(t *testing.T) {
	t.Parallel()

	t.Run("drop database", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "DROP DATABASE customers"
			expected := &ast.DropDatabaseStatement{
				Database: "customers",
			}

			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			require.NoError(t, err)
			assert.Equal(t, expected, stmts)
		})

		t.Run("wrong database name", func(t *testing.T) {
			t.Parallel()

			input := "DROP DATABASE 9"
			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			require.NotNil(t, err)
			assert.Nil(t, stmts)
		})
	})

	t.Run("drop table", func(t *testing.T) {
		t.Parallel()

		t.Run("correct query", func(t *testing.T) {
			t.Parallel()

			input := "DROP TABLE customers"
			expected := &ast.DropTableStatement{
				Table: "customers",
			}

			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			require.NoError(t, err)
			assert.Equal(t, expected, stmts)
		})

		t.Run("wrong table name", func(t *testing.T) {
			t.Parallel()

			input := "DROP TABLE +"
			p := parser.New(lexer.New(input))
			stmts, err := p.Parse()

			require.NotNil(t, err)
			assert.Nil(t, stmts)
		})
	})

	t.Run("drop unexpected", func(t *testing.T) {
		t.Parallel()

		input := "DROP abc"
		p := parser.New(lexer.New(input))
		stmts, err := p.Parse()

		require.NotNil(t, err)
		assert.Nil(t, stmts)
	})
}

func TestParser_Parse(t *testing.T) {
	t.Parallel()

	t.Run("returns nil if input is empty", func(t *testing.T) {
		t.Parallel()

		p := parser.New(lexer.New(""))
		stmt, err := p.Parse()

		require.Nil(t, err)
		assert.Nil(t, stmt)
	})
}
