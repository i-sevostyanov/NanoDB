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
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestDelete_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)
		rows := []sql.Row{
			{datatype.NewInteger(1), datatype.NewString("Max")},
			{datatype.NewInteger(2), datatype.NewString("Vlad")},
			{datatype.NewInteger(3), datatype.NewString("John")},
		}

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		deleter := NewMockRowDeleter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),

			rowIter.EXPECT().Next().Return(rows[0], nil),
			deleter.EXPECT().Delete(int64(1)).Return(nil),

			rowIter.EXPECT().Next().Return(rows[1], nil),
			deleter.EXPECT().Delete(int64(2)).Return(nil),

			rowIter.EXPECT().Next().Return(rows[2], nil),
			deleter.EXPECT().Delete(int64(3)).Return(nil),

			rowIter.EXPECT().Next().Return(nil, io.EOF),
			rowIter.EXPECT().Close().Return(nil),
		)

		deletePlan := plan.NewDelete(deleter, pkIndex, child)
		iter, err := deletePlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		assert.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")

		child := plan.NewMockNode(ctrl)
		deleter := NewMockRowDeleter(ctrl)

		child.EXPECT().RowIter().Return(nil, expectedErr)

		deletePlan := plan.NewDelete(deleter, pkIndex, child)
		iter, err := deletePlan.RowIter()
		require.NotNil(t, err)
		assert.Nil(t, iter)
	})

	t.Run("return error on Next call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		deleter := NewMockRowDeleter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(nil, expectedErr),
			rowIter.EXPECT().Close().Return(nil),
		)

		deletePlan := plan.NewDelete(deleter, pkIndex, child)
		iter, err := deletePlan.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on delete call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")

		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewString("Max"),
		}

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		deleter := NewMockRowDeleter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			deleter.EXPECT().Delete(int64(1)).Return(expectedErr),
			rowIter.EXPECT().Close().Return(nil),
		)

		deletePlan := plan.NewDelete(deleter, pkIndex, child)
		iter, err := deletePlan.RowIter()
		require.NoError(t, err)

		r, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, r)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on unsupported key type", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)
		row := sql.Row{
			datatype.NewString("Max"),
		}

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		deleter := NewMockRowDeleter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			rowIter.EXPECT().Close().Return(nil),
		)

		deletePlan := plan.NewDelete(deleter, pkIndex, child)
		iter, err := deletePlan.RowIter()
		require.NoError(t, err)

		r, err := iter.Next()
		require.NotNil(t, err)
		assert.Nil(t, r)

		err = iter.Close()
		require.NoError(t, err)
	})
}
