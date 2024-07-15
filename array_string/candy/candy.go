package candy

func candy(ratings []int) int {
	var candies []int

	for range ratings {
		candies = append(candies, 1)
	}

	// check from left to right
	// the LEFT element must have more candies
	for i := 1; i < len(ratings); i++ {
		j := i - 1
		if ratings[i] > ratings[j] {
			candies[i] = candies[j] + 1
		}
	}

	// check from right to left
	// the RIGHT element must have more candies
	for i := len(ratings) - 2; i >= 0; i-- {
		j := i + 1
		if ratings[i] > ratings[j] {
			candies[i] = max(candies[i], candies[j]+1)
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}

	return sum
}
