package rotate_array

func rotate(nums []int, k int) {
	// rotateOn(nums, k)
	rotateO1(nums, k)
}

func rotateOn(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}

	tmp := make([]int, len(nums))

	for i, num := range nums {
		j := (i + k) % len(nums)
		tmp[j] = num
	}

	copy(nums, tmp)
}

func rotateO1(nums []int, k int) {
	k %= len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
