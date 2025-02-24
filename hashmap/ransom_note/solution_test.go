package ransom_note

import (
	"testing"
)

func Test_canConstruct(t *testing.T) {
	res := canConstruct("a", "b")
	t.Log(res)

	res = canConstruct("aa", "ab")
	t.Log(res)

	res = canConstruct("aa", "aab")
	t.Log(res)

	res = canConstruct("aa", "aaba")
	t.Log(res)
}
