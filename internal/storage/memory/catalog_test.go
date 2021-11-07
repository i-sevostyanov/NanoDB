package memory_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/storage/memory"
)

func TestCatalog_GetDatabase(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		name := "playground"
		catalog := memory.NewCatalog()
		expected, err := catalog.CreateDatabase(name)
		require.NoError(t, err)

		database, err := catalog.GetDatabase(name)
		require.NoError(t, err)
		require.Equal(t, expected, database)
	})

	t.Run("returns error if database not exist", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()
		database, err := catalog.GetDatabase("not-exist")
		require.NotNil(t, err)
		require.Nil(t, database)
	})
}

func TestCatalog_ListDatabases(t *testing.T) {
	t.Parallel()

	t.Run("returns empty list", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()
		databases, err := catalog.ListDatabases()
		require.NoError(t, err)
		require.Empty(t, databases)
	})

	t.Run("returns databases without errors", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()

		first, err := catalog.CreateDatabase("first")
		require.NoError(t, err)

		second, err := catalog.CreateDatabase("second")
		require.NoError(t, err)

		expected := []sql.Database{
			first,
			second,
		}

		databases, err := catalog.ListDatabases()
		require.NoError(t, err)
		require.ElementsMatch(t, expected, databases)
	})
}

func TestCatalog_CreateDatabase(t *testing.T) {
	t.Parallel()

	t.Run("create database successfully", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()
		db, err := catalog.CreateDatabase("playground")
		require.NoError(t, err)
		require.NotNil(t, db)
	})

	t.Run("returns error if database already exist", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()
		db, err := catalog.CreateDatabase("playground")
		require.NoError(t, err)
		require.NotNil(t, db)

		db, err = catalog.CreateDatabase("playground")
		require.NotNil(t, err)
		require.Nil(t, db)
	})
}

func TestCatalog_DropDatabase(t *testing.T) {
	t.Parallel()

	t.Run("drop table without errors", func(t *testing.T) {
		t.Parallel()

		database := "playground"

		catalog := memory.NewCatalog()
		db, err := catalog.CreateDatabase(database)
		require.NoError(t, err)
		require.NotNil(t, db)

		err = catalog.DropDatabase(database)
		require.NoError(t, err)

		_, err = catalog.GetDatabase(database)
		require.NotNil(t, err)
	})

	t.Run("returns error if database not exist", func(t *testing.T) {
		t.Parallel()

		catalog := memory.NewCatalog()
		err := catalog.DropDatabase("playground")
		require.NotNil(t, err)
	})
}
