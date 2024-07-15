package remove_duplicates_from_sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// expected := 5
	res := removeDuplicates(nums)

	assert.Equal(t, expected, res)
}
