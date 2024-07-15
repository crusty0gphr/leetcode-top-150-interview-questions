package text_justification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	in := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	expected := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	actual := fullJustify(in, maxWidth)
	assert.Equal(t, expected, actual)

	in = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	expected = []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	actual = fullJustify(in, maxWidth)
	assert.Equal(t, expected, actual)

	in = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	expected = []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}
	actual = fullJustify(in, maxWidth)
	assert.Equal(t, expected, actual)
}
