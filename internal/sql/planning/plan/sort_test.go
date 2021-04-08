package plan_test

import (
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestSort_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("sorting integers", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewInteger(1)},
			{datatype.NewInteger(3)},
			{datatype.NewInteger(4)},
			{datatype.NewInteger(2)},
			{datatype.NewInteger(5)},
		}

		t.Run("ascending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
				rowIter.EXPECT().Close().Return(nil),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.ErrorIs(t, err, io.EOF)
			assert.Nil(t, row)

			err = iter.Close()
			assert.NoError(t, err)
		})

		t.Run("descending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)
		})
	})

	t.Run("sorting strings", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewString("Anna")},
			{datatype.NewString("Catalina")},
			{datatype.NewString("Bob")},
			{datatype.NewString("Eval")},
			{datatype.NewString("David")},
		}

		t.Run("ascending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)
		})

		t.Run("descending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)
		})
	})

	t.Run("sorting floats", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewFloat(1)},
			{datatype.NewFloat(3)},
			{datatype.NewFloat(4)},
			{datatype.NewFloat(2)},
			{datatype.NewFloat(5)},
		}

		t.Run("ascending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)
		})

		t.Run("descending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)
		})
	})

	t.Run("sorting booleans", func(t *testing.T) {
		t.Parallel()

		isTrue := sql.Row{datatype.NewBoolean(true)}
		isFalse := sql.Row{datatype.NewBoolean(false)}

		rows := []sql.Row{
			isTrue,
			isFalse,
			isTrue,
			isFalse,
			isTrue,
		}

		t.Run("ascending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)
		})

		t.Run("descending", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)
		})
	})

	t.Run("sorting integer/null", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewInteger(1)},
			{datatype.NewNull()},
			{datatype.NewInteger(4)},
			{datatype.NewNull()},
			{datatype.NewInteger(5)},
		}

		t.Run("ascending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)
		})

		t.Run("descending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)
		})
	})

	t.Run("sorting float/null", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewFloat(1)},
			{datatype.NewNull()},
			{datatype.NewFloat(4)},
			{datatype.NewNull()},
			{datatype.NewFloat(5)},
		}

		t.Run("ascending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)
		})

		t.Run("descending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)
		})
	})

	t.Run("sorting string/null", func(t *testing.T) {
		t.Parallel()

		rows := []sql.Row{
			{datatype.NewString("1")},
			{datatype.NewNull()},
			{datatype.NewString("4")},
			{datatype.NewNull()},
			{datatype.NewString("5")},
		}

		t.Run("ascending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)
		})

		t.Run("descending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[4], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[2], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[0], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[1], row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, rows[3], row)
		})
	})

	t.Run("sorting boolean/null", func(t *testing.T) {
		t.Parallel()

		isTrue := sql.Row{datatype.NewBoolean(true)}
		isFalse := sql.Row{datatype.NewBoolean(false)}
		null := sql.Row{datatype.NewNull()}

		rows := []sql.Row{
			isFalse,
			null,
			isTrue,
			null,
			isFalse,
		}

		t.Run("ascending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Ascending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, null, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, null, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)
		})

		t.Run("descending", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			order := plan.Descending
			columnPos := uint8(0)

			child := plan.NewMockNode(ctrl)
			rowIter := sql.NewMockRowIter(ctrl)

			gomock.InOrder(
				child.EXPECT().RowIter().Return(rowIter, nil),
				rowIter.EXPECT().Next().Return(rows[0], nil),
				rowIter.EXPECT().Next().Return(rows[1], nil),
				rowIter.EXPECT().Next().Return(rows[2], nil),
				rowIter.EXPECT().Next().Return(rows[3], nil),
				rowIter.EXPECT().Next().Return(rows[4], nil),
				rowIter.EXPECT().Next().Return(nil, io.EOF),
			)

			sort := plan.NewSort(columnPos, order, child)
			iter, err := sort.RowIter()
			require.NoError(t, err)
			require.NotNil(t, iter)

			row, err := iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isTrue, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, isFalse, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, null, row)

			row, err = iter.Next()
			require.NoError(t, err)
			assert.Equal(t, null, row)
		})
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		order := plan.Descending
		columnPos := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")

		child := plan.NewMockNode(ctrl)
		child.EXPECT().RowIter().Return(nil, expectedErr)

		sort := plan.NewSort(columnPos, order, child)
		iter, err := sort.RowIter()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})

	t.Run("returns error on expected sort order", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		order := plan.Order(math.MaxUint8)
		columnPos := uint8(0)

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		child.EXPECT().RowIter().Return(rowIter, nil)

		sort := plan.NewSort(columnPos, order, child)
		iter, err := sort.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.NotNil(t, err)
		require.Nil(t, row)
	})

	t.Run("returns error on Next call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		order := plan.Descending
		columnPos := uint8(0)
		expectedErr := fmt.Errorf("something went wrong")

		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(nil, expectedErr),
		)

		sort := plan.NewSort(columnPos, order, child)
		iter, err := sort.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		assert.Nil(t, row)
	})
}
