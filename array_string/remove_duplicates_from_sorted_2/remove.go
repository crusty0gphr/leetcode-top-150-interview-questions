package remove_duplicates_from_sorted

// pairs = 2
//
// {0, 0, 1, 1, 1, 1, 2, 3, 3} -> if i == j-2
//                    i
//              j
//       j-2

func removeDuplicates(nums []int) int {
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
