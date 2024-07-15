package find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	left, right := 0, len(needle)-1

	for right < len(haystack) {
		window := haystack[left : right+1]
		if window == needle {
			return left
		}

		left++
		right++
	}

	return -1
}
