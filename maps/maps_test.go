package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	ks := Keys(map[int]float32{
		1: 1.5,
		2: 3.5,
	})

	assert.ElementsMatch(t, ks, []int{1, 2})
}
