package valid_anagram

func isAnagram(s string, t string) bool {
	check := make(map[rune]int, len(s))

	for _, ch := range s {
		check[ch]++
	}

	for _, ch := range t {
		check[ch]--
		if check[ch] == 0 {
			delete(check, ch)
		}
	}

	return len(check) == 0
}
