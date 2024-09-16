package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	actual := maxArea(height)
	assert.Equal(t, expected, actual)
}
