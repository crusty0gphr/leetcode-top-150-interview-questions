package sum

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var res [][]int
	const target = 0

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := num + nums[left] + nums[right]

			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				res = append(res, []int{num, nums[left], nums[right]})
				left++

				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}

	return res
}
