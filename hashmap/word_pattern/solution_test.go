package word_pattern

import (
	"testing"
)

func Test_wordPattern(t *testing.T) {
	res := wordPattern("abba", "dog cat cat dog")
	t.Log(res)

	res = wordPattern("abba", "dog cat cat parrot")
	t.Log(res)
}
