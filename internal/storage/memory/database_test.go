package memory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/storage/memory"
)

func TestDatabase_ListTables(t *testing.T) {
	t.Parallel()

	t.Run("returns empty list", func(t *testing.T) {
		t.Parallel()

		database := memory.NewDatabase("playground")
		tables := database.ListTables()
		require.Empty(t, tables)
	})

	t.Run("returns tables", func(t *testing.T) {
		t.Parallel()

		database := memory.NewDatabase("playground")
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

		users, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		tickets, err := database.CreateTable("tickets", scheme)
		require.NoError(t, err)

		expected := []sql.Table{
			users,
			tickets,
		}

		tables := database.ListTables()
		require.ElementsMatch(t, expected, tables)
	})
}

func TestDatabase_GetTable(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
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
		expected, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		table, err := database.GetTable("users")
		require.NoError(t, err)
		require.Equal(t, expected, table)
	})

	t.Run("returns error if table not exist", func(t *testing.T) {
		t.Parallel()

		database := memory.NewDatabase("playground")
		table, err := database.GetTable("xxx")
		require.NotNil(t, err)
		require.Nil(t, table)
	})
}

func TestDatabase_CreateTable(t *testing.T) {
	t.Parallel()

	t.Run("create table without errors", func(t *testing.T) {
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
	})

	t.Run("returns error if table already exist", func(t *testing.T) {
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
		_, err := database.CreateTable("users", scheme)
		require.NoError(t, err)

		_, err = database.CreateTable("users", scheme)
		require.NotNil(t, err)
	})

	t.Run("return error if scheme is empty", func(t *testing.T) {
		t.Parallel()

		database := memory.NewDatabase("playground")
		_, err := database.CreateTable("users", nil)
		require.NotNil(t, err)
	})
}

func TestDatabase_DropTable(t *testing.T) {
	t.Parallel()

	t.Run("drop table without errors", func(t *testing.T) {
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

		err = database.DropTable(table.Name())
		require.NoError(t, err)

		deleted, err := database.GetTable(table.Name())
		require.NotNil(t, err)
		require.Nil(t, deleted)
	})

	t.Run("returns error if table not exist", func(t *testing.T) {
		t.Parallel()

		database := memory.NewDatabase("playground")
		err := database.DropTable("xxx")
		require.NotNil(t, err)
	})
}

func TestDatabase_Name(t *testing.T) {
	t.Parallel()

	name := "playground"
	database := memory.NewDatabase(name)
	assert.Equal(t, name, database.Name())
}
