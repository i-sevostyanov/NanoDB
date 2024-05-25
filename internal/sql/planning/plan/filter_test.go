package plan_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestFilter_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	columns := []string{"id", "name"}

	child := plan.NewMockNode(ctrl)
	cond := expr.NewMockNode(ctrl)

	child.EXPECT().Columns().Return(columns)

	filter := plan.NewFilter(cond, child)
	assert.Equal(t, columns, filter.Columns())
}

func TestFilter_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cond := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		rows := []sql.Row{
			{datatype.NewInteger(1), datatype.NewText("Max")},
			{datatype.NewInteger(2), datatype.NewText("Vlad")},
			{datatype.NewInteger(3), datatype.NewText("John")},
		}

		isTrue := datatype.NewBoolean(true)
		isFalse := datatype.NewBoolean(false)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),

			rowIter.EXPECT().Next().Return(rows[0], nil),
			cond.EXPECT().Eval(rows[0]).Return(isTrue, nil),

			rowIter.EXPECT().Next().Return(rows[1], nil),
			cond.EXPECT().Eval(rows[1]).Return(isFalse, nil),

			rowIter.EXPECT().Next().Return(rows[2], nil),
			cond.EXPECT().Eval(rows[2]).Return(isFalse, nil),

			rowIter.EXPECT().Next().Return(nil, io.EOF),
			rowIter.EXPECT().Close().Return(nil),
		)

		filter := plan.NewFilter(cond, child)
		iter, err := filter.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, rows[0], row)

		row, err = iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("something went wrong")

		cond := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)

		child.EXPECT().RowIter().Return(nil, expectedErr)

		filter := plan.NewFilter(cond, child)
		iter, err := filter.RowIter()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})

	t.Run("returns error on Next call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("something went wrong")

		cond := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(nil, expectedErr),
			rowIter.EXPECT().Close().Return(nil),
		)

		filter := plan.NewFilter(cond, child)
		iter, err := filter.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on Eval call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cond := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		expectedErr := fmt.Errorf("something went wrong")
		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
		}

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			cond.EXPECT().Eval(row).Return(nil, expectedErr),
			rowIter.EXPECT().Close().Return(nil),
		)

		filter := plan.NewFilter(cond, child)
		iter, err := filter.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err = iter.Next()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on unexpected condition value", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cond := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		value := sql.NewMockValue(ctrl)

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
		}

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			cond.EXPECT().Eval(row).Return(value, nil),
			value.EXPECT().Raw().Return(10),
			rowIter.EXPECT().Close().Return(nil),
		)

		filter := plan.NewFilter(cond, child)
		iter, err := filter.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err = iter.Next()
		require.Error(t, err)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})
}
