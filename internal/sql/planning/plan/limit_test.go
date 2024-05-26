package plan_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestLimit_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	columns := []string{"id", "name"}

	child := plan.NewMockNode(ctrl)
	child.EXPECT().Columns().Return(columns)

	limit := plan.NewLimit(10, child)
	assert.Equal(t, columns, limit.Columns())
}

func TestLimit_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("negative limit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)

		limit := plan.NewLimit(-10, child)
		iter, err := limit.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		assert.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("zero limit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)

		limit := plan.NewLimit(0, child)
		iter, err := limit.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		assert.Equal(t, io.EOF, err)
		assert.Nil(t, row)
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rows := []sql.Row{
			{datatype.NewInteger(1)},
			{datatype.NewInteger(2)},
			{datatype.NewInteger(3)},
		}

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(rows[0], nil),
			rowIter.EXPECT().Next().Return(rows[1], nil),
			rowIter.EXPECT().Close().Return(nil),
		)

		limit := plan.NewLimit(2, child)
		iter, err := limit.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.NoError(t, err)
		assert.Equal(t, rows[0], row)

		row, err = iter.Next()
		require.NoError(t, err)
		assert.Equal(t, rows[1], row)

		row, err = iter.Next()
		assert.Equal(t, io.EOF, err)
		assert.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("somethig went wtrong")

		child := plan.NewMockNode(ctrl)
		child.EXPECT().RowIter().Return(nil, expectedErr)

		limit := plan.NewLimit(2, child)
		iter, err := limit.RowIter()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})

	t.Run("returns error on Next call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("somethig went wtrong")

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)
		rowIter.EXPECT().Next().Return(nil, expectedErr)

		limit := plan.NewLimit(2, child)
		iter, err := limit.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, row)
	})
}
