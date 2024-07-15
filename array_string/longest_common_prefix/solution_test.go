package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	expected := "fl"
	actual := longestCommonPrefix(input)
	assert.Equal(t, expected, actual)

	input = []string{"dog", "racecar", "car"}
	expected = ""
	actual = longestCommonPrefix(input)
	assert.Equal(t, expected, actual)

	input = []string{"dog"}
	expected = "dog"
	actual = longestCommonPrefix(input)
	assert.Equal(t, expected, actual)

	input = []string{"aaaaaaa", "aaa", "aaaaa", "aaaaaaaa"}
	expected = "aaa"
	actual = longestCommonPrefix(input)
	assert.Equal(t, expected, actual)
}
