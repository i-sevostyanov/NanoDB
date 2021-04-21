package memory_test

import (
	"testing"

	"github.com/i-sevostyanov/NanoDB/internal/storage/memory"
	"github.com/stretchr/testify/assert"
)

func TestSequence_Next(t *testing.T) {
	t.Parallel()

	seq := &memory.Sequence{}
	assert.Equal(t, int64(0), seq.Value())
	assert.Equal(t, int64(1), seq.Next())
	assert.Equal(t, int64(2), seq.Next())
	assert.Equal(t, int64(3), seq.Next())
}

func TestSequence_SetValue(t *testing.T) {
	t.Parallel()

	seq := &memory.Sequence{}

	seq.SetValue(99)
	assert.Equal(t, int64(99), seq.Value())

	seq.Next()
	assert.Equal(t, int64(100), seq.Value())
}

func TestSequence_Value(t *testing.T) {
	t.Parallel()

	seq := &memory.Sequence{}
	assert.Equal(t, int64(0), seq.Value())
	seq.Next()
	seq.Next()
	seq.Next()
	assert.Equal(t, int64(3), seq.Value())
}
