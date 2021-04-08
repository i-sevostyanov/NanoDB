package plan_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i-sevostyanov/NanoDB/internal/sql/planning/plan"
)

func TestEmpty_RowIter(t *testing.T) {
	t.Parallel()

	emptyPlan := plan.NewEmpty()
	iter, err := emptyPlan.RowIter()
	require.NoError(t, err)
	assert.NotNil(t, iter)

	row, err := iter.Next()
	require.Equal(t, io.EOF, err)
	assert.Nil(t, row)

	err = iter.Close()
	require.NoError(t, err)
}
