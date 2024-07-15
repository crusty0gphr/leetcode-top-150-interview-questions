package jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := jump(nums)
	assert.Equal(t, 2, res)
}
