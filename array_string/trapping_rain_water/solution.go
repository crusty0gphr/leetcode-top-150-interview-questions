package trapping_rain_water

// 0  1  0  2  1  0  1  3  2  1  2  1
// l
//                                  r
// -----------------------------------
// maxL = 1 | maxR = 1
// v = v + maxL - h[l] | v = v + maxR - h[R]
// -----------------------------------
// maxL < maxR ? l++, r--
// -----------------------------------
// maxL < maxR ? maxL = max(maxL & h[l]) : maxR = (max maxR & h[r])

func trap(height []int) int {
	waterVolume := 0
	if len(height) == 0 {
		return waterVolume
	}

	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			waterVolume += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			waterVolume += maxRight - height[right]
		}
	}

	return waterVolume
}
