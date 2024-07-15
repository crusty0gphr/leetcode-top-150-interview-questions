package length_of_last_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	input := "Hello World"
	expected := 5
	actual := lengthOfLastWord(input)
	assert.Equal(t, expected, actual)

	input = "   fly me   to   the moon  "
	expected = 4
	actual = lengthOfLastWord(input)
	assert.Equal(t, expected, actual)

	input = "luffy is still joyboy"
	expected = 6
	actual = lengthOfLastWord(input)
	assert.Equal(t, expected, actual)

	input = "a"
	expected = 1
	actual = lengthOfLastWord(input)
	assert.Equal(t, expected, actual)
}
