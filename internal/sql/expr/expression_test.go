package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
)

func TestNewExpr(t *testing.T) {
	t.Parallel()

	t.Run("column expr", func(t *testing.T) {
		t.Parallel()

		t.Run("returns column expr", func(t *testing.T) {
			t.Parallel()

			column := "name"
			scheme := sql.Scheme{
				column: sql.Column{
					Position:   1,
					Name:       "name",
					DataType:   sql.String,
					PrimaryKey: false,
					Nullable:   false,
					Default:    nil,
				},
			}

			expected := expr.Column{
				Name:     "name",
				Position: scheme[column].Position,
			}

			identExpr := &ast.IdentExpr{
				Name: column,
			}

			node, err := expr.New(identExpr, scheme)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("returns error if column not exists", func(t *testing.T) {
			t.Parallel()

			scheme := sql.Scheme{
				"name": sql.Column{
					Position:   1,
					Name:       "name",
					DataType:   sql.String,
					PrimaryKey: false,
					Nullable:   false,
					Default:    nil,
				},
			}

			identExpr := &ast.IdentExpr{
				Name: "id",
			}

			node, err := expr.New(identExpr, scheme)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})

		t.Run("returns error on empty scheme", func(t *testing.T) {
			t.Parallel()

			identExpr := &ast.IdentExpr{
				Name: "name",
			}

			node, err := expr.New(identExpr, nil)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})
	})

	t.Run("unary expr", func(t *testing.T) {
		t.Parallel()

		t.Run("unary plus", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.UnaryExpr{
				Operator: token.Add,
				Right: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
			}

			operand, err := expr.NewInteger("10")
			require.NoError(t, err)

			expected := &expr.Unary{
				Operator: expr.UnaryPlus,
				Operand:  operand,
			}

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)

			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("unary minus", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.UnaryExpr{
				Operator: token.Sub,
				Right: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
			}

			operand, err := expr.NewInteger("10")
			require.NoError(t, err)

			expected := &expr.Unary{
				Operator: expr.UnaryMinus,
				Operand:  operand,
			}

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("unexpected unary operator", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.UnaryExpr{
				Operator: token.Illegal,
				Right: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
			}

			node, err := expr.New(astExpr, nil)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})

		t.Run("unexpected scalar type on the right", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.UnaryExpr{
				Operator: token.Sub,
				Right: &ast.ScalarExpr{
					Type:    token.Illegal,
					Literal: "10",
				},
			}

			node, err := expr.New(astExpr, nil)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})
	})

	t.Run("scalar expr", func(t *testing.T) {
		t.Parallel()

		t.Run("integer", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.ScalarExpr{
				Type:    token.Integer,
				Literal: "10",
			}

			expected, err := expr.NewInteger("10")
			require.NoError(t, err)

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("float", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.ScalarExpr{
				Type:    token.Float,
				Literal: "10",
			}

			expected, err := expr.NewFloat("10")
			require.NoError(t, err)

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("string", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.ScalarExpr{
				Type:    token.String,
				Literal: "xyz",
			}

			expected, err := expr.NewString("xyz")
			require.NoError(t, err)

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("boolean", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.ScalarExpr{
				Type:    token.Boolean,
				Literal: "true",
			}

			expected, err := expr.NewBoolean("true")
			require.NoError(t, err)

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("null", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.ScalarExpr{
				Type:    token.Null,
				Literal: "null",
			}

			expected := expr.Null{}

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})
	})

	t.Run("binary expr", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.BinaryExpr{
				Left: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
				Operator: token.Add,
				Right: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "100",
				},
			}

			left, err := expr.NewInteger("10")
			require.NoError(t, err)

			right, err := expr.NewInteger("100")
			require.NoError(t, err)

			expected := &expr.Binary{
				Operator: expr.Add,
				Left:     left,
				Right:    right,
			}

			node, err := expr.New(astExpr, nil)
			require.NoError(t, err)
			assert.Equal(t, expected, node)
		})

		t.Run("unexpected scalar type on the left", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.BinaryExpr{
				Left: &ast.ScalarExpr{
					Type:    token.Illegal,
					Literal: "10",
				},
				Operator: token.Add,
				Right: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "100",
				},
			}

			node, err := expr.New(astExpr, nil)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})

		t.Run("unexpected scalar type on the right", func(t *testing.T) {
			t.Parallel()

			astExpr := &ast.BinaryExpr{
				Left: &ast.ScalarExpr{
					Type:    token.Integer,
					Literal: "10",
				},
				Operator: token.Add,
				Right: &ast.ScalarExpr{
					Type:    token.Illegal,
					Literal: "100",
				},
			}

			node, err := expr.New(astExpr, nil)
			require.NotNil(t, err)
			assert.Nil(t, node)
		})
	})

	t.Run("return error on unexpected expression type", func(t *testing.T) {
		t.Parallel()

		astExpr := &ast.AsteriskExpr{}
		node, err := expr.New(astExpr, nil)
		require.NotNil(t, err)
		assert.Nil(t, node)
	})
}
