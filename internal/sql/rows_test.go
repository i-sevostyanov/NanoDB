package sql_test

import (
	"io"
	"testing"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/stretchr/testify/require"
)

func TestSliceRowsIter_Next(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		iter := sql.RowsIter()
		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("one row", func(t *testing.T) {
		t.Parallel()

		expected := sql.Row{
			datatype.NewInteger(10),
		}

		iter := sql.RowsIter(expected)
		row, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, expected, row)

		row, err = iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)
	})

	t.Run("multiple rows", func(t *testing.T) {
		t.Parallel()

		row1 := sql.Row{datatype.NewInteger(1)}
		row2 := sql.Row{datatype.NewInteger(2)}
		row3 := sql.Row{datatype.NewInteger(3)}
		iter := sql.RowsIter(row1, row2, row3)

		row, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, row1, row)

		row, err = iter.Next()
		require.NoError(t, err)
		require.Equal(t, row2, row)

		row, err = iter.Next()
		require.NoError(t, err)
		require.Equal(t, row3, row)

		row, err = iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)
	})
}
