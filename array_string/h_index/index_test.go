package h_index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	input := []int{3, 0, 6, 1, 5}
	hIdx := 3

	res := hIndex(input)
	assert.Equal(t, hIdx, res)
}
