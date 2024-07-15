package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3

	expected := []int{1, 2, 2, 3, 5, 6}

	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, expected)
}
