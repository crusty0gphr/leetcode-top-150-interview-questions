package trapping_rain_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6

	res := trap(height)
	assert.Equal(t, expected, res)

	height = []int{4, 2, 0, 3, 2, 5}
	expected = 9

	res = trap(height)
	assert.Equal(t, expected, res)
}
