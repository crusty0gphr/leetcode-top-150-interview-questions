package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2

	res := removeElement(nums, val)
	assert.Equal(t, expected, res)
}
