package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	in := "A man, a plan, a canal: Panama"
	expected := true
	actual := isPalindrome(in)

	assert.Equal(t, expected, actual)

	in = "race a car"
	expected = false
	actual = isPalindrome(in)

	assert.Equal(t, expected, actual)

	in = " "
	expected = true
	actual = isPalindrome(in)

	assert.Equal(t, expected, actual)
}
