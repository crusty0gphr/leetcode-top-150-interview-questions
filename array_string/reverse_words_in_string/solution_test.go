package reverse_words_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	expected := "blue is sky the"

	actual := reverseWords(s)
	assert.Equal(t, expected, actual)

	s = "  hello world  "
	expected = "world hello"

	actual = reverseWords(s)
	assert.Equal(t, expected, actual)

	s = "a good   example"
	expected = "example good a"

	actual = reverseWords(s)
	assert.Equal(t, expected, actual)
}
