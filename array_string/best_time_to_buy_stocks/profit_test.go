package best_time_to_buy_stocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	res := maxProfit(prices)

	assert.Equal(t, expected, res)
}
