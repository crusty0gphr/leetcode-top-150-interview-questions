package find_the_index_of_the_first_occurrence_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	expected := 0
	actual := strStr(haystack, needle)
	assert.Equal(t, expected, actual)

	haystack = "leetcode"
	needle = "leeto"
	expected = -1
	actual = strStr(haystack, needle)
	assert.Equal(t, expected, actual)
}
