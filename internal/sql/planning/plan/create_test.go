package plan_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestCreateDatabase_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	creator := NewMockDatabaseCreator(ctrl)
	createPlan := plan.NewCreateDatabase(creator, "test")
	assert.Nil(t, createPlan.Columns())
}

func TestCreateDatabase_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		creator := NewMockDatabaseCreator(ctrl)
		creator.EXPECT().CreateDatabase(name).Return(nil, nil)

		createPlan := plan.NewCreateDatabase(creator, name)
		iter, err := createPlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on database creation", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		expectedErr := fmt.Errorf("something went wrong")

		creator := NewMockDatabaseCreator(ctrl)
		creator.EXPECT().CreateDatabase(name).Return(nil, expectedErr)

		createPlan := plan.NewCreateDatabase(creator, name)
		iter, err := createPlan.RowIter()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})
}

func TestCreateTable_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	creator := NewMockTableCreator(ctrl)
	createPlan := plan.NewCreateTable(creator, "test", nil)
	assert.Nil(t, createPlan.Columns())
}

func TestCreateTable_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "users"
		scheme := sql.Scheme{}

		creator := NewMockTableCreator(ctrl)
		creator.EXPECT().CreateTable(name, scheme).Return(nil, nil)

		createPlan := plan.NewCreateTable(creator, name, scheme)
		iter, err := createPlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on table creation", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "test"
		scheme := sql.Scheme{}
		expectedErr := fmt.Errorf("something went wrong")

		creator := NewMockTableCreator(ctrl)
		creator.EXPECT().CreateTable(name, scheme).Return(nil, expectedErr)

		createPlan := plan.NewCreateTable(creator, name, scheme)
		iter, err := createPlan.RowIter()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})
}
