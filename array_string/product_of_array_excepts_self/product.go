package product_of_array_excepts_self

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix, suffix := 1, 1

	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
