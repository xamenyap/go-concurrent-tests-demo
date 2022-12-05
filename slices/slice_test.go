package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	// calling t.Parallel() to mark this test as being eligible for concurrent run.
	t.Parallel()
	got := Unique([]string{"foo", "bar"})
	assert.True(t, got)
}

func TestExists(t *testing.T) {
	t.Parallel()

	t.Run("strings", func(t *testing.T) {
		// t.Parallel() has to be called inside each sub test
		// to fully benefit from Go concurrency.
		t.Parallel()
		got := Exists([]string{"foo", "bar"}, "bar")
		assert.True(t, got)
	})

	t.Run("integers", func(t *testing.T) {
		// t.Parallel() has to be called inside each sub test
		// to fully benefit from Go concurrency.
		t.Parallel()
		got := Exists([]int{1, 2, 3}, 2)
		assert.True(t, got)
	})
}

func TestIndex(t *testing.T) {
	// will not run in parallel with TestExists and TestUnique
	idx := Index([]string{"foo", "bar"}, "bar")
	assert.Equal(t, 1, idx)
}
