package gas_station

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	target := 3 // id

	res := canCompleteCircuit(gas, cost)
	assert.Equal(t, target, res)

	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	target = -1 // id

	res = canCompleteCircuit(gas, cost)
	assert.Equal(t, target, res)
}
