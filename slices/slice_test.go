package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		got := Unique([]int{1, 2, 3})
		assert.True(t, got)
	})

	t.Run("strings", func(t *testing.T) {
		got := Unique([]string{"foo", "bar"})
		assert.True(t, got)
	})
}

func TestExists(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		got := Exists([]int{1, 2, 3}, 2)
		assert.True(t, got)
	})

	t.Run("strings", func(t *testing.T) {
		got := Exists([]string{"foo", "bar"}, "bar")
		assert.True(t, got)
	})
}
