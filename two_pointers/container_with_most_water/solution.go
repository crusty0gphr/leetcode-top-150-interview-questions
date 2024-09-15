package container_with_most_water

func maxArea(heights []int) int {
	res := 0
	left, right := 0, len(heights)-1

	for left < right {
		height := min(heights[left], heights[right])
		width := right - left
		area := height * width

		res = max(res, area)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
