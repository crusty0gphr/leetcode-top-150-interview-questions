package isomorphic_string

func isIsomorphic(s string, t string) bool {
	s2tMap := [128]uint8{}
	t2sMap := [128]uint8{}

	for i := 0; i < len(s); i++ {
		charInS := int(s[i])
		charInT := int(t[i])

		if s2tMap[charInS] == 0 && t2sMap[charInT] == 0 {
			s2tMap[charInS] = t[i]
			t2sMap[charInT] = s[i]
		}

		if s2tMap[charInS] != t[i] || t2sMap[charInT] != s[i] {
			return false
		}
	}

	return true
}
