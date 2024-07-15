package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := canJump(nums)
	assert.Equal(t, true, res)

	nums = []int{3, 2, 1, 0, 4}
	res = canJump(nums)
	assert.Equal(t, false, res)
}
