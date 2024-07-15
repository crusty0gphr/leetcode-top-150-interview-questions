package jump_game

// [  2  3  1  1  4  ]
//    l
//       r
// 1 jump

func jump(nums []int) int {
	res := 0
	left, right := 0, 0

	for right < len(nums)-1 {
		farthest := 0
		for i := left; i < right+1; i++ {
			farthest = max(farthest, i+nums[i])
		}
		left = right + 1
		right = farthest
		res++
	}
	return res
}
