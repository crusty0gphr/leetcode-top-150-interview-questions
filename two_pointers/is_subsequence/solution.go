package is_subsequence

func isSubsequence(s string, t string) bool {
	cursor := 0

	for i := 0; i < len(t); i++ {
		if cursor < len(s) && t[i] == s[cursor] {
			cursor++
		}
	}
	return cursor == len(s)
}
