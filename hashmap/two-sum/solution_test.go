package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	assert.Equal(t, res, []int{0, 1})

	nums = []int{3, 2, 4}
	target = 6

	res = twoSum(nums, target)
	assert.Equal(t, res, []int{1, 2})
}
