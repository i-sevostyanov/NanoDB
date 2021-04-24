package expr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql"
	"github.com/i-sevostyanov/NanoDB/internal/sql/datatype"
	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

func TestBinary_String(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	leftNode := expr.NewMockNode(ctrl)
	rightNode := expr.NewMockNode(ctrl)

	operator := expr.And
	expected := fmt.Sprintf("(left %s right)", operator.String())
	binaryExpr := expr.Binary{
		Operator: operator,
		Left:     leftNode,
		Right:    rightNode,
	}

	leftNode.EXPECT().String().Return("left")
	rightNode.EXPECT().String().Return("right")

	value := binaryExpr.String()
	assert.Equal(t, expected, value)
}

func TestBinary_Eval(t *testing.T) {
	t.Parallel()

	t.Run("equal", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Equal,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Equal(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Equal,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Equal(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("not equal", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.NotEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().NotEqual(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.NotEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().NotEqual(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("greater than", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.GreaterThan,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().GreaterThan(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.GreaterThan,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().GreaterThan(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("less than", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.LessThan,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().LessThan(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.LessThan,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().LessThan(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("greater or equal", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.GreaterThanOrEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().GreaterOrEqual(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.GreaterThanOrEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().GreaterOrEqual(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("less or equal", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.LessThanOrEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().LessOrEqual(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.LessThanOrEqual,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().LessOrEqual(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("and", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.And,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().And(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.And,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().And(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("or", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Or,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Or(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Or,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Or(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("add", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Add,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Add(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Add,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Add(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("sub", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Sub,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Sub(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Sub,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Sub(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("mul", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Mul,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Mul(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Mul,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Mul(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("div", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Div,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Div(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Div,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Div(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("mod", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Mod,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Mod(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Mod,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Mod(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("pow", func(t *testing.T) {
		t.Parallel()

		t.Run("no error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Pow,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Pow(rightValue).Return(expected, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.NoError(t, err)
			assert.Equal(t, expected, value)
		})

		t.Run("eval returns error", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			leftNode := expr.NewMockNode(ctrl)
			rightNode := expr.NewMockNode(ctrl)

			binaryExpr := expr.Binary{
				Operator: expr.Pow,
				Left:     leftNode,
				Right:    rightNode,
			}

			leftValue := sql.NewMockValue(ctrl)
			rightValue := sql.NewMockValue(ctrl)
			expected := fmt.Errorf("unexpected error")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
				leftValue.EXPECT().Pow(rightValue).Return(nil, expected),
			)

			value, err := binaryExpr.Eval(nil)
			assert.Equal(t, expected, err)
			assert.Nil(t, value)
		})
	})

	t.Run("left node returns error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		leftNode := expr.NewMockNode(ctrl)
		rightNode := expr.NewMockNode(ctrl)

		equal := expr.Binary{
			Operator: expr.Equal,
			Left:     leftNode,
			Right:    rightNode,
		}

		expected := fmt.Errorf("unexpected error")
		leftNode.EXPECT().Eval(nil).Return(nil, expected)

		value, err := equal.Eval(nil)
		require.True(t, errors.Is(err, expected))
		assert.Nil(t, value)
	})

	t.Run("right node returns error", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		leftNode := expr.NewMockNode(ctrl)
		rightNode := expr.NewMockNode(ctrl)

		binaryExpr := expr.Binary{
			Operator: expr.Equal,
			Left:     leftNode,
			Right:    rightNode,
		}

		expected := fmt.Errorf("unexpected error")
		leftValue := sql.NewMockValue(ctrl)

		gomock.InOrder(
			leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
			rightNode.EXPECT().Eval(nil).Return(nil, expected),
		)

		value, err := binaryExpr.Eval(nil)
		require.True(t, errors.Is(err, expected))
		assert.Nil(t, value)
	})

	t.Run("returns error on unexpected operator", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		leftNode := expr.NewMockNode(ctrl)
		rightNode := expr.NewMockNode(ctrl)

		binaryExpr := expr.Binary{
			Operator: expr.BinaryOp("p"),
			Left:     leftNode,
			Right:    rightNode,
		}

		leftValue := sql.NewMockValue(ctrl)
		rightValue := sql.NewMockValue(ctrl)

		gomock.InOrder(
			leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
			rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
		)

		value, err := binaryExpr.Eval(nil)
		require.NotNil(t, err)
		assert.Nil(t, value)
	})
}
