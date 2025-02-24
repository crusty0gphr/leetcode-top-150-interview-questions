package word_pattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	w2pMap := make(map[string]uint8)
	p2wMap := make(map[uint8]string)

	for i := 0; i < len(words); i++ {
		char := pattern[i]
		word := words[i]

		_, pExists := w2pMap[word]
		_, wExists := p2wMap[char]

		if !pExists && !wExists {
			w2pMap[word] = char
			p2wMap[char] = word
		}

		if w2pMap[word] != char || p2wMap[char] != word {
			return false
		}
	}

	return true
}
