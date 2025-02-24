package valid_anagram

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	res := isAnagram("anagram", "nagaram")
	fmt.Println(res)

	res = isAnagram("rat", "cat")
	fmt.Println(res)

	res = isAnagram("a", "ab")
	fmt.Println(res)

	res = isAnagram("ab", "a")
	fmt.Println(res)
}
