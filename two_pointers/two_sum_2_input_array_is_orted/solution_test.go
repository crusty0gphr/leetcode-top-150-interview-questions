package two_sum_2_input_array_is_orted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	in := []int{2, 7, 11, 15}
	target := 9
	expected := []int{1, 2}

	actual := twoSum(in, target)
	assert.Equal(t, expected, actual)

	in = []int{2, 3, 4}
	target = 6
	expected = []int{1, 3}

	actual = twoSum(in, target)
	assert.Equal(t, expected, actual)

	in = []int{-1, 0}
	target = -1
	expected = []int{1, 2}

	actual = twoSum(in, target)
	assert.Equal(t, expected, actual)

	in = []int{5, 25, 75}
	target = 100
	expected = []int{2, 3}

	actual = twoSum(in, target)
	assert.Equal(t, expected, actual)
}
