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

func TestUpdate_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	columns := []string{"id", "name"}

	updater := NewMockRowUpdater(ctrl)
	child := plan.NewMockNode(ctrl)
	child.EXPECT().Columns().Return(columns)

	update := plan.NewUpdate(updater, 1, nil, child)
	assert.Equal(t, columns, update.Columns())
}

func TestUpdate_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		pkIndex := uint8(0)

		nameExpr := expr.NewMockNode(ctrl)
		salaryExpr := expr.NewMockNode(ctrl)

		columns := map[uint8]expr.Node{
			1: nameExpr,
			2: salaryExpr,
		}

		firstID := sql.NewMockValue(ctrl)
		secondID := sql.NewMockValue(ctrl)
		thirdID := sql.NewMockValue(ctrl)

		rows := []sql.Row{
			{firstID, sql.NewMockValue(ctrl), sql.NewMockValue(ctrl)},
			{secondID, sql.NewMockValue(ctrl), sql.NewMockValue(ctrl)},
			{thirdID, sql.NewMockValue(ctrl), sql.NewMockValue(ctrl)},
		}

		updated := []sql.Row{
			{firstID, datatype.NewText("Max"), datatype.NewFloat(2000)},
			{secondID, datatype.NewText("Jane"), datatype.NewFloat(2200)},
			{thirdID, datatype.NewText("John"), datatype.NewFloat(2100)},
		}

		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)

		rowIter.EXPECT().Next().Return(rows[0], nil)
		nameExpr.EXPECT().Eval(rows[0]).Return(updated[0][1], nil)
		salaryExpr.EXPECT().Eval(rows[0]).Return(updated[0][2], nil)
		firstID.EXPECT().Raw().Return(int64(1))
		updater.EXPECT().Update(int64(1), updated[0]).Return(nil)

		rowIter.EXPECT().Next().Return(rows[1], nil)
		nameExpr.EXPECT().Eval(rows[1]).Return(updated[1][1], nil)
		salaryExpr.EXPECT().Eval(rows[1]).Return(updated[1][2], nil)
		secondID.EXPECT().Raw().Return(int64(2))
		updater.EXPECT().Update(int64(2), updated[1]).Return(nil)

		rowIter.EXPECT().Next().Return(rows[2], nil)
		nameExpr.EXPECT().Eval(rows[2]).Return(updated[2][1], nil)
		salaryExpr.EXPECT().Eval(rows[2]).Return(updated[2][2], nil)
		thirdID.EXPECT().Raw().Return(int64(3))
		updater.EXPECT().Update(int64(3), updated[2]).Return(nil)

		rowIter.EXPECT().Next().Return(nil, io.EOF)
		rowIter.EXPECT().Close().Return(nil)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.NoError(t, err)

		row, err := iter.Next()
		require.ErrorIs(t, err, io.EOF)
		assert.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")
		columns := map[uint8]expr.Node{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)

		child.EXPECT().RowIter().Return(nil, expectedErr)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, iter)
	})

	t.Run("returns error on Next call", func(t *testing.T) {
		t.Parallel()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")
		columns := map[uint8]expr.Node{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)
		rowIter.EXPECT().Next().Return(nil, expectedErr)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, row)
	})

	t.Run("returns error on eval call", func(t *testing.T) {
		t.Parallel()

		pkIndex := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")
		row := sql.Row{
			datatype.NewInteger(1),
			datatype.NewText("Max"),
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		nameExpr := expr.NewMockNode(ctrl)

		columns := map[uint8]expr.Node{
			1: nameExpr,
		}

		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)
		rowIter.EXPECT().Next().Return(row, nil)
		nameExpr.EXPECT().Eval(row).Return(nil, expectedErr)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err = iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, row)
	})

	t.Run("returns error on unsupported key type", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		nameExpr := expr.NewMockNode(ctrl)
		id := sql.NewMockValue(ctrl)
		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		pkIndex := uint8(0)
		columns := map[uint8]expr.Node{
			1: nameExpr,
		}

		row := sql.Row{
			id,
			datatype.NewText("Max"),
		}

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			nameExpr.EXPECT().Eval(row).Return(nil, nil),
			id.EXPECT().Raw().Return("xyz"),
		)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err = iter.Next()
		require.NotNil(t, err)
		assert.Nil(t, row)
	})

	t.Run("returns error on update", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		nameExpr := expr.NewMockNode(ctrl)
		id := sql.NewMockValue(ctrl)
		updater := NewMockRowUpdater(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		expectedErr := fmt.Errorf("something went wrong")
		key := int64(1)
		pkIndex := uint8(0)
		columns := map[uint8]expr.Node{
			1: nameExpr,
		}

		row := sql.Row{
			id,
			datatype.NewText("Max"),
		}

		updated := sql.Row{
			id,
			datatype.NewText("John"),
		}

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(row, nil),
			nameExpr.EXPECT().Eval(row).Return(updated[1], nil),
			id.EXPECT().Raw().Return(key),
			updater.EXPECT().Update(key, updated).Return(expectedErr),
		)

		update := plan.NewUpdate(updater, pkIndex, columns, child)
		iter, err := update.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err = iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, row)
	})
}
