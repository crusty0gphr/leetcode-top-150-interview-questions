package two_sum

func twoSum(nums []int, target int) []int {
	check := make(map[int]int, len(nums))

	for i, num := range nums {
		if val, exists := check[target-num]; exists {
			return []int{i, val}
		}

		check[num] = i
	}

	return nil
}
