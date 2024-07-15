package remove_duplicates_from_sorted

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	// nums := []int{1, 1, 1, 2, 2, 3}
	// expected := 5
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expected := 7
	res := removeDuplicates(nums)

	fmt.Println(nums)

	assert.Equal(t, expected, res)
}
