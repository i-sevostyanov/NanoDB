package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql/expr"
)

func TestInteger(t *testing.T) {
	t.Parallel()

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		integer, err := expr.NewInteger("10")
		require.NoError(t, err)
		assert.Equal(t, "10", integer.String())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewInteger("10")
		require.NoError(t, err)
	})

	t.Run("returns error on unexpected input", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewInteger("xyz")
		require.Error(t, err)
	})

	t.Run("eval", func(t *testing.T) {
		t.Parallel()

		integer, err := expr.NewInteger("10")
		require.NoError(t, err)

		value, err := integer.Eval(nil)
		require.NoError(t, err)

		switch v := value.Raw().(type) {
		case int64:
			assert.Equal(t, int64(10), v)
		default:
			assert.Failf(t, "fail", "unexpected value type %T", v)
		}
	})
}

func TestFloat(t *testing.T) {
	t.Parallel()

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		float, err := expr.NewFloat("10")
		require.NoError(t, err)
		assert.Equal(t, "10", float.String())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewFloat("10")
		require.NoError(t, err)
	})

	t.Run("returns error on unexpected input", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewFloat("xyz")
		require.Error(t, err)
	})

	t.Run("eval", func(t *testing.T) {
		t.Parallel()

		float, err := expr.NewFloat("10")
		require.NoError(t, err)

		value, err := float.Eval(nil)
		require.NoError(t, err)

		switch v := value.Raw().(type) {
		case float64:
			assert.InEpsilon(t, float64(10), v, 0)
		default:
			assert.Failf(t, "fail", "unexpected value type %T", v)
		}
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		str, err := expr.NewString("xyz")
		require.NoError(t, err)
		assert.Equal(t, "xyz", str.String())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewString("xyz")
		require.NoError(t, err)
	})

	t.Run("eval", func(t *testing.T) {
		t.Parallel()

		s, err := expr.NewString("xyz")
		require.NoError(t, err)

		value, err := s.Eval(nil)
		require.NoError(t, err)

		switch v := value.Raw().(type) {
		case string:
			assert.Equal(t, "xyz", v)
		default:
			assert.Failf(t, "fail", "unexpected value type %T", v)
		}
	})
}

func TestBoolean(t *testing.T) {
	t.Parallel()

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		boolean, err := expr.NewBoolean("true")
		require.NoError(t, err)
		assert.Equal(t, "true", boolean.String())
	})

	t.Run("no error", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewBoolean("true")
		require.NoError(t, err)
	})

	t.Run("returns error on unexpected input", func(t *testing.T) {
		t.Parallel()

		_, err := expr.NewBoolean("xyz")
		require.Error(t, err)
	})

	t.Run("eval", func(t *testing.T) {
		t.Parallel()

		boolean, err := expr.NewBoolean("true")
		require.NoError(t, err)

		value, err := boolean.Eval(nil)
		require.NoError(t, err)

		switch v := value.Raw().(type) {
		case bool:
			assert.True(t, v)
		default:
			assert.Failf(t, "fail", "unexpected value type %T", v)
		}
	})
}

func TestNull(t *testing.T) {
	t.Parallel()

	t.Run("string()", func(t *testing.T) {
		t.Parallel()

		null := expr.NewNull()
		assert.Equal(t, "null", null.String())
	})

	t.Run("eval", func(t *testing.T) {
		t.Parallel()

		null := expr.NewNull()
		value, err := null.Eval(nil)
		require.NoError(t, err)

		switch v := value.Raw().(type) {
		case nil:
		default:
			assert.Failf(t, "fail", "unexpected value type %T", v)
		}
	})
}
