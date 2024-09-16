package two_sum_2_input_array_is_orted

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		actual := numbers[left] + numbers[right]

		if actual > target {
			right--
		}

		if actual < target {
			left++
		}

		if actual == target {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}
