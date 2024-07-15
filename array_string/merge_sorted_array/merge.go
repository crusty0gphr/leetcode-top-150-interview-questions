package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1 // last elements
	j := n - 1 // last elements

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
}
