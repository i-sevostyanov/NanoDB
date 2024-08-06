package expr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

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

			leftValue := datatype.NewText("123")
			rightValue := datatype.NewText("456")
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewBoolean(true)
			rightValue := datatype.NewInteger(10)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewText("123")
			rightValue := datatype.NewText("123")
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewText("123")
			rightValue := datatype.NewFloat(123)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(100)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewText("123")
			rightValue := datatype.NewFloat(123)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(100)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewText("123")
			rightValue := datatype.NewFloat(123)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(100)
			rightValue := datatype.NewFloat(100)
			expected := datatype.NewBoolean(true)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(100)
			rightValue := datatype.NewText("100")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(1000)
			rightValue := datatype.NewFloat(100)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(100)
			rightValue := datatype.NewText("100")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewBoolean(true)
			rightValue := datatype.NewBoolean(false)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewBoolean(true)
			rightValue := datatype.NewText("true")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewBoolean(false)
			rightValue := datatype.NewBoolean(false)
			expected := datatype.NewBoolean(false)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewBoolean(false)
			rightValue := datatype.NewText("false")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewFloat(20)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewText("10")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewFloat(0)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewText("10")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewFloat(100)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewText("10")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewFloat(1)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(0)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(10)
			expected := datatype.NewFloat(0)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewText("10")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewFloat(2)
			expected := datatype.NewFloat(100)

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
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

			leftValue := datatype.NewFloat(10)
			rightValue := datatype.NewText("10")

			gomock.InOrder(
				leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
				rightNode.EXPECT().Eval(nil).Return(rightValue, nil),
			)

			value, err := binaryExpr.Eval(nil)
			require.Error(t, err)
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

		expected := errors.New("unexpected error")
		leftNode.EXPECT().Eval(nil).Return(nil, expected)

		value, err := equal.Eval(nil)
		require.ErrorIs(t, err, expected)
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

		expected := errors.New("unexpected error")
		leftValue := sql.NewMockValue(ctrl)

		gomock.InOrder(
			leftNode.EXPECT().Eval(nil).Return(leftValue, nil),
			rightNode.EXPECT().Eval(nil).Return(nil, expected),
		)

		value, err := binaryExpr.Eval(nil)
		require.ErrorIs(t, err, expected)
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
		require.Error(t, err)
		assert.Nil(t, value)
	})
}
