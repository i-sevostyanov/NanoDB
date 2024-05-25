package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

func TestColumn_String(t *testing.T) {
	t.Parallel()

	name := "id"
	column := expr.Column{
		Name:     name,
		Position: 1,
	}

	assert.Equal(t, name, column.String())
}

func TestColumn_Eval(t *testing.T) {
	t.Parallel()

	t.Run("returns name from row by index", func(t *testing.T) {
		t.Parallel()

		expected := "John Doe"
		column := expr.Column{Position: 1}

		// Table scheme:
		//  id
		//  name
		//  salary
		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText(expected),
			datatype.NewFloat(1500.5),
		}

		value, err := column.Eval(row)
		require.NoError(t, err)
		assert.Equal(t, expected, value.Raw())
	})

	t.Run("returns error on empty row", func(t *testing.T) {
		t.Parallel()

		column := expr.Column{Position: 1}
		value, err := column.Eval(nil)

		require.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("returns error if position out of range", func(t *testing.T) {
		t.Parallel()

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
			datatype.NewFloat(1500.5),
		}

		column := expr.Column{Position: uint8(len(row) + 1)}
		value, err := column.Eval(row)

		require.Error(t, err)
		assert.Nil(t, value)
	})
}
