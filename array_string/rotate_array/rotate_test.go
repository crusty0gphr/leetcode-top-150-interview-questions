package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, k)
	assert.Equal(t, expected, nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	expected = []int{3, 99, -1, -100}
	rotate(nums, k)
	assert.Equal(t, expected, nums)

	nums = []int{1, 2, 3, 4}
	k = 4
	expected = []int{1, 2, 3, 4}
	rotate(nums, k)
	assert.Equal(t, expected, nums)

	// nums := []int{1}
	// k := 4
	// expected := []int{1}
	// rotate(nums, k)
	// assert.Equal(t, expected, nums)
}
