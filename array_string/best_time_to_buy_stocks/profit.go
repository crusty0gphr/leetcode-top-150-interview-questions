package best_time_to_buy_stocks

func maxProfit(prices []int) int {
	profit := 0
	left, right := 0, 1

	for right < len(prices) {
		if prices[right]-prices[left] > profit {
			profit = prices[right] - prices[left]
		}

		if prices[left] > prices[right] {
			left++
		} else {
			right++
		}
	}
	return profit
}
