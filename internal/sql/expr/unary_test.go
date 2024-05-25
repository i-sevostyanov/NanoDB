package expr_test

import (
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

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		t.Run("unary plus", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := "(operand)"

			operand := expr.NewMockNode(ctrl)
			operand.EXPECT().String().Return("operand")

			unaryExpr := expr.Unary{
				Operator: expr.UnaryPlus,
				Operand:  operand,
			}

			value := unaryExpr.String()
			assert.Equal(t, expected, value)
		})

		t.Run("unary minus", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := "(-operand)"

			operand := expr.NewMockNode(ctrl)
			operand.EXPECT().String().Return("operand")

			unaryExpr := expr.Unary{
				Operator: expr.UnaryMinus,
				Operand:  operand,
			}

			value := unaryExpr.String()
			assert.Equal(t, expected, value)
		})
	})

	t.Run("unary plus", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := expr.NewMockNode(ctrl)
		sqlValue := sql.NewMockValue(ctrl)

		expected := datatype.NewInteger(10)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryPlus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(sqlValue, nil)
		sqlValue.EXPECT().DataType().Return(expected.DataType())
		sqlValue.EXPECT().Raw().Return(expected.Raw())

		value, err := unaryExpr.Eval(nil)
		require.NoError(t, err)
		assert.Equal(t, expected, value)
	})

	t.Run("unary minus", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := expr.NewMockNode(ctrl)
		sqlValue := sql.NewMockValue(ctrl)

		expected := datatype.NewInteger(-10)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryMinus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(sqlValue, nil)
		sqlValue.EXPECT().DataType().Return(expected.DataType())
		sqlValue.EXPECT().Raw().Return(int64(10))

		value, err := unaryExpr.Eval(nil)
		require.NoError(t, err)
		assert.Equal(t, expected, value)
	})

	t.Run("unexpected unary operator", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := expr.NewMockNode(ctrl)
		unaryExpr := expr.Unary{
			Operator: expr.UnaryOp(math.MaxUint64),
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(nil, nil)

		value, err := unaryExpr.Eval(nil)
		require.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("operand returns error on eval", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		operand := expr.NewMockNode(ctrl)
		expected := fmt.Errorf("something went wrong")
		unaryExpr := expr.Unary{
			Operator: expr.UnaryMinus,
			Operand:  operand,
		}

		operand.EXPECT().Eval(nil).Return(nil, expected)

		value, err := unaryExpr.Eval(nil)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, value)
	})
}
