package plan_test

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestInsert_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	key := int64(1)
	row := sql.Row{
		datatype.NewInteger(key),
		datatype.NewText("Max"),
	}

	inserter := NewMockTableInserter(ctrl)
	insertPlan := plan.NewInsert(inserter, key, row)
	assert.Nil(t, insertPlan.Columns())
}

func TestInsert_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		key := int64(1)
		row := sql.Row{
			datatype.NewInteger(key),
			datatype.NewText("Max"),
		}

		inserter := NewMockTableInserter(ctrl)
		inserter.EXPECT().Insert(key, row).Return(nil)

		insertPlan := plan.NewInsert(inserter, key, row)
		iter, err := insertPlan.RowIter()
		require.NoError(t, err)

		row, err = iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on insert row", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		key := int64(1)
		expectedErr := errors.New("something went wrong")
		row := sql.Row{
			datatype.NewInteger(key),
			datatype.NewText("Max"),
		}

		inserter := NewMockTableInserter(ctrl)
		inserter.EXPECT().Insert(key, row).Return(expectedErr)

		createPlan := plan.NewInsert(inserter, key, row)
		iter, err := createPlan.RowIter()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})
}
