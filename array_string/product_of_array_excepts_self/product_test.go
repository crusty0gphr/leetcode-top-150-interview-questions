package product_of_array_excepts_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	res := productExceptSelf(input)
	assert.Equal(t, expected, res)

	input = []int{-1, 1, 0, -3, 3}
	expected = []int{0, 0, 9, 0, 0}

	res = productExceptSelf(input)
	assert.Equal(t, expected, res)
}
