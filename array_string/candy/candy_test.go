package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy(t *testing.T) {
	rating := []int{5, 4, 3, 5, 6, 2}
	target := 12

	res := candy(rating)
	assert.Equal(t, target, res)
}
