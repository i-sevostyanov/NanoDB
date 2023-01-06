package memory_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/storage/memory"
)

func TestTable_Name(t *testing.T) {
	t.Parallel()

	tableName := "users"
	scheme := sql.Scheme{
		"id": sql.Column{
			Position:   0,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
	}

	database := memory.NewDatabase("playground")
	table, err := database.CreateTable(tableName, scheme)
	require.NoError(t, err)
	assert.Equal(t, tableName, table.Name())
}

func TestTable_Scheme(t *testing.T) {
	t.Parallel()

	scheme := sql.Scheme{
		"id": sql.Column{
			Position:   0,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
	}

	database := memory.NewDatabase("playground")
	table, err := database.CreateTable("users", scheme)
	require.NoError(t, err)
	assert.Equal(t, scheme, table.Scheme())
}

func TestTable_PrimaryKey(t *testing.T) {
	t.Parallel()

	scheme := sql.Scheme{
		"name": sql.Column{
			Position:   0,
			Name:       "name",
			DataType:   sql.Text,
			PrimaryKey: false,
			Nullable:   false,
			Default:    nil,
		},
		"id": sql.Column{
			Position:   1,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
	}

	database := memory.NewDatabase("playground")
	table, err := database.CreateTable("users", scheme)
	require.NoError(t, err)
	assert.Equal(t, scheme["id"], table.PrimaryKey())
}

func TestTable_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"name": sql.Column{
				Position:   0,
				Name:       "name",
				DataType:   sql.Text,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		iter, err := table.Scan()
		require.NoError(t, err)

		row, err := iter.Next()
		require.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})
}

func TestTable_Sequence(t *testing.T) {
	t.Parallel()

	scheme := sql.Scheme{
		"id": sql.Column{
			Position:   0,
			Name:       "id",
			DataType:   sql.Integer,
			PrimaryKey: true,
			Nullable:   false,
			Default:    nil,
		},
	}

	database := memory.NewDatabase("playground")
	table, err := database.CreateTable("users", scheme)
	require.NoError(t, err)
	assert.NotNil(t, table.Sequence())
}

func TestTable_Insert(t *testing.T) {
	t.Parallel()

	t.Run("insert without errors", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"name": sql.Column{
				Position:   0,
				Name:       "name",
				DataType:   sql.Text,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		key := int64(1)
		expected := sql.Row{
			datatype.NewText("Max"),
			datatype.NewInteger(key),
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Insert(key, expected)
		require.NoError(t, err)

		iter, err := table.Scan()
		require.NoError(t, err)

		row, err := iter.Next()
		require.NoError(t, err)
		assert.Equal(t, expected, row)

		row, err = iter.Next()
		require.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on duplicate key", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		key := int64(1)
		expected := sql.Row{
			datatype.NewInteger(key),
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Insert(key, expected)
		require.NoError(t, err)

		err = table.Insert(key, nil)
		require.NotNil(t, err)
	})
}

func TestTable_Delete(t *testing.T) {
	t.Parallel()

	t.Run("deletes one row", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		key := int64(1)
		expected := sql.Row{
			datatype.NewInteger(key),
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Insert(key, expected)
		require.NoError(t, err)

		err = table.Delete(key)
		require.NoError(t, err)

		iter, err := table.Scan()
		require.NoError(t, err)

		row, err := iter.Next()
		require.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("deletes all rows", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		rows := []struct {
			key int64
			row sql.Row
		}{
			{
				key: 1,
				row: sql.Row{datatype.NewInteger(1)},
			},
			{
				key: 2,
				row: sql.Row{datatype.NewInteger(2)},
			},
			{
				key: 3,
				row: sql.Row{datatype.NewInteger(3)},
			},
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		for _, r := range rows {
			err = table.Insert(r.key, r.row)
			require.NoError(t, err)
		}

		for _, r := range rows {
			err = table.Delete(r.key)
			require.NoError(t, err)
		}

		iter, err := table.Scan()
		require.NoError(t, err)

		row, err := iter.Next()
		require.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error if key not found", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Delete(1)
		require.NotNil(t, err)
	})
}

func TestTable_Update(t *testing.T) {
	t.Parallel()

	t.Run("update without errors", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   0,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
			"name": sql.Column{
				Position:   1,
				Name:       "name",
				DataType:   sql.Text,
				PrimaryKey: false,
				Nullable:   false,
				Default:    nil,
			},
		}

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
		}

		updated := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Vlad"),
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Insert(1, row)
		require.NoError(t, err)

		err = table.Update(1, updated)
		require.NoError(t, err)

		iter, err := table.Scan()
		require.NoError(t, err)

		actual, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, updated, actual)

		row, err = iter.Next()
		require.ErrorIs(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error if key not found", func(t *testing.T) {
		t.Parallel()

		scheme := sql.Scheme{
			"id": sql.Column{
				Position:   1,
				Name:       "id",
				DataType:   sql.Integer,
				PrimaryKey: true,
				Nullable:   false,
				Default:    nil,
			},
		}

		database := memory.NewDatabase("playground")
		table, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		err = table.Update(1, nil)
		require.NotNil(t, err)
	})
}
