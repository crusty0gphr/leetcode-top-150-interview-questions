package remove_duplicates_from_sorted

// {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
//  i     j

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			nums = append(nums[:i], nums[j:]...)
		} else {
			i = j
			j++
		}
	}
	return len(nums)
}
