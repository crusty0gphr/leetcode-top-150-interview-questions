package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	target := "ahbgdc"
	str := "abc"
	expected := true
	actual := isSubsequence(str, target)

	assert.Equal(t, expected, actual)

	target = "ahbgdc"
	str = ""
	expected = true
	actual = isSubsequence(str, target)

	assert.Equal(t, expected, actual)

	target = "abc"
	str = "b"
	expected = true
	actual = isSubsequence(str, target)

	assert.Equal(t, expected, actual)
}
