package expr_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

func TestUnary_Eval(t *testing.T) {
	t.Parallel()

	t.Run("unary plus", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := NewMockNode(ctrl)
		sqlValue := sql.NewMockValue(ctrl)

		expected := datatype.NewInteger(10)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryPlus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(sqlValue, nil)
		sqlValue.EXPECT().UnaryPlus().Return(expected, nil)

		value, err := unaryExpr.Eval(nil)
		require.NoError(t, err)
		assert.Equal(t, expected, value)
	})

	t.Run("unary minus", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := NewMockNode(ctrl)
		sqlValue := sql.NewMockValue(ctrl)

		expected := datatype.NewInteger(-10)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryMinus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(sqlValue, nil)
		sqlValue.EXPECT().UnaryMinus().Return(expected, nil)

		value, err := unaryExpr.Eval(nil)
		require.NoError(t, err)
		assert.Equal(t, expected, value)
	})

	t.Run("unexpected unary operator", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := NewMockNode(ctrl)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryOp(math.MaxUint64),
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(nil, nil)

		value, err := unaryExpr.Eval(nil)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})

	t.Run("operand returns error on eval", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := NewMockNode(ctrl)
		expected := fmt.Errorf("something went wrong")
		unaryExpr := expr.Unary{
			Operator: expr.UnaryMinus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(nil, expected)

		value, err := unaryExpr.Eval(nil)
		require.True(t, errors.Is(err, expected))
		assert.Nil(t, value)
	})
}
