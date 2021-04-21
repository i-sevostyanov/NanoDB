package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

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
			datatype.NewString(expected),
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

		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("returns error if position out of range", func(t *testing.T) {
		t.Parallel()

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewString("Max"),
			datatype.NewFloat(1500.5),
		}

		column := expr.Column{Position: uint8(len(row) + 1)}
		value, err := column.Eval(row)

		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}
