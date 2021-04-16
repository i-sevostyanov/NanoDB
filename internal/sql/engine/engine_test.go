package engine_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/engine"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/ast"
	"github.com/i-sevostyanov/NanoDB/internal/sql/parsing/token"
	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
	"github.com/stretchr/testify/require"
)

func TestEngine_Query(t *testing.T) {
	t.Parallel()

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		input := "select true"
		database := "playground"
		astNode := &ast.SelectStatement{
			Result: []ast.ResultStatement{
				{
					Alias: nil,
					Expr: &ast.ScalarExpr{
						Type:    token.Boolean,
						Literal: "true",
					},
				},
			},
		}

		parser := NewMockParser(ctrl)
		planner := NewMockPlanner(ctrl)
		rowIter := sql.NewMockRowIter(ctrl)
		planNode := plan.NewMockNode(ctrl)

		parser.EXPECT().Parse(input).Return(astNode, nil)
		planner.EXPECT().Plan(database, astNode).Return(planNode, nil)
		planNode.EXPECT().RowIter().Return(rowIter, nil)

		ng := engine.New(parser, planner)
		iter, err := ng.Exec(database, input)
		require.NoError(t, err)
		require.NotNil(t, iter)
	})

	t.Run("returns an error if the parse fails", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		input := "select true"
		database := "playground"
		expectedErr := fmt.Errorf("something went wrong")

		parser := NewMockParser(ctrl)
		planner := NewMockPlanner(ctrl)

		parser.EXPECT().Parse(input).Return(nil, expectedErr)

		ng := engine.New(parser, planner)
		iter, err := ng.Exec(database, input)
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})

	t.Run("returns an error if the planner fails", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		input := "select true"
		database := "playground"
		expectedErr := fmt.Errorf("something went wrong")
		astNode := &ast.SelectStatement{
			Result: []ast.ResultStatement{
				{
					Alias: nil,
					Expr: &ast.ScalarExpr{
						Type:    token.Boolean,
						Literal: "true",
					},
				},
			},
		}

		parser := NewMockParser(ctrl)
		planner := NewMockPlanner(ctrl)

		parser.EXPECT().Parse(input).Return(astNode, nil)
		planner.EXPECT().Plan(database, astNode).Return(nil, expectedErr)

		ng := engine.New(parser, planner)
		iter, err := ng.Exec(database, input)
		require.ErrorIs(t, err, expectedErr)
		require.Nil(t, iter)
	})
}
