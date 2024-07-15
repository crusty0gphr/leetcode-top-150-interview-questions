package length_of_last_word

func lengthOfLastWord(s string) int {
	var length int
	prev, last := len(s)-1, len(s)-1

	for {
		if s[last] == ' ' {
			last--
			prev--
		} else {
			length++
			prev--
			if prev < 0 || s[prev] == ' ' {
				return length
			}
		}
	}
}
