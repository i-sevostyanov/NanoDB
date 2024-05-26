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
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestProject_Columns(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	child := plan.NewMockNode(ctrl)
	id := expr.NewMockNode(ctrl)
	name := expr.NewMockNode(ctrl)

	id.EXPECT().String().Return("id")

	projections := []plan.Projection{
		{
			Expr: id,
		},
		{
			Alias: "X",
			Expr:  name,
		},
	}

	columns := []string{
		"id",
		"X",
	}

	project := plan.NewProject(projections, child)
	assert.Equal(t, columns, project.Columns())
}

func TestProject_RowIter(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rows := []sql.Row{
			{datatype.NewInteger(1), datatype.NewText("Greg"), datatype.NewFloat(2000)},
			{datatype.NewInteger(2), datatype.NewText("Frank"), datatype.NewFloat(2200)},
		}

		expectedRows := []sql.Row{
			{datatype.NewInteger(1), datatype.NewText("Greg")},
			{datatype.NewInteger(2), datatype.NewText("Frank")},
		}

		id := expr.NewMockNode(ctrl)
		name := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),

			rowIter.EXPECT().Next().Return(rows[0], nil),
			id.EXPECT().Eval(rows[0]).Return(rows[0][0], nil),
			name.EXPECT().Eval(rows[0]).Return(rows[0][1], nil),

			rowIter.EXPECT().Next().Return(rows[1], nil),
			id.EXPECT().Eval(rows[1]).Return(rows[1][0], nil),
			name.EXPECT().Eval(rows[1]).Return(rows[1][1], nil),

			rowIter.EXPECT().Next().Return(nil, io.EOF),
			rowIter.EXPECT().Close().Return(nil),
		)

		projections := []plan.Projection{
			{
				Expr: id,
			},
			{
				Expr: name,
			},
		}

		project := plan.NewProject(projections, child)
		iter, err := project.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, expectedRows[0], row)

		row, err = iter.Next()
		require.NoError(t, err)
		require.Equal(t, expectedRows[1], row)

		row, err = iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)

		err = iter.Close()
		require.NoError(t, err)
	})

	t.Run("returns nothing", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := expr.NewMockNode(ctrl)
		name := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(nil, io.EOF),
		)

		projections := []plan.Projection{
			{
				Expr: id,
			},
			{
				Expr: name,
			},
		}

		project := plan.NewProject(projections, child)
		iter, err := project.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.Equal(t, io.EOF, err)
		require.Nil(t, row)
	})

	t.Run("returns error on eval", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rows := []sql.Row{
			{datatype.NewInteger(1), datatype.NewText("Greg"), datatype.NewFloat(2000)},
			{datatype.NewInteger(2), datatype.NewText("Frank"), datatype.NewFloat(2200)},
		}

		expectedErr := fmt.Errorf("something went wrong")

		id := expr.NewMockNode(ctrl)
		name := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(rows[0], nil),
			id.EXPECT().Eval(rows[0]).Return(nil, expectedErr),
		)

		projections := []plan.Projection{
			{
				Expr: id,
			},
			{
				Expr: name,
			},
		}

		project := plan.NewProject(projections, child)
		iter, err := project.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, row)
	})

	t.Run("returns error on RowIter call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("something went wrong")

		id := expr.NewMockNode(ctrl)
		name := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)

		child.EXPECT().RowIter().Return(nil, expectedErr)

		projections := []plan.Projection{
			{
				Expr: id,
			},
			{
				Expr: name,
			},
		}

		project := plan.NewProject(projections, child)
		iter, err := project.RowIter()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})

	t.Run("returns error on next call", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedErr := fmt.Errorf("something went wrong")

		id := expr.NewMockNode(ctrl)
		name := expr.NewMockNode(ctrl)
		child := plan.NewMockNode(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)

		gomock.InOrder(
			child.EXPECT().RowIter().Return(rowIter, nil),
			rowIter.EXPECT().Next().Return(nil, expectedErr),
		)

		projections := []plan.Projection{
			{
				Expr: id,
			},
			{
				Expr: name,
			},
		}

		project := plan.NewProject(projections, child)
		iter, err := project.RowIter()
		require.NoError(t, err)
		require.NotNil(t, iter)

		row, err := iter.Next()
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, row)
	})
}
