package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	abc := [26]int{}

	for _, char := range magazine {
		abc[char-'a']++
	}

	for _, char := range ransomNote {
		abc[char-'a']--
		if abc[char-'a'] < 0 {
			return false
		}
	}

	return true
}
