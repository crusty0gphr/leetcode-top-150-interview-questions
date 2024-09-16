package valid_palindrome

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	var buff []rune

	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			continue
		}
		buff = append(buff, unicode.ToLower(char))
	}

	fmt.Println(string(buff))
	left, right := 0, len(buff)-1
	for left < right {
		if buff[left] != buff[right] {
			return false
		}

		left++
		right--
	}

	return true
}
