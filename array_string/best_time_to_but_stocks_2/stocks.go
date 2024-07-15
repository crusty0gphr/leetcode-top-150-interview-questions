package best_time_to_but_stocks_2

// 7, 1, 5, 3, 6, 4
//    l, r

func maxProfit(prices []int) int {
	profit := 0
	left, right := 0, 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit += prices[right] - prices[left]
		}
		left++
		right++
	}

	return profit
}
