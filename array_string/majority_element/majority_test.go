package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2
	// nums := []int{3, 2, 3}
	// expected := 3

	res := majorityElement(nums)
	assert.Equal(t, expected, res)
}
