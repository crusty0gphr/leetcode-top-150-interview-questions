package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	num := 3749
	expected := "MMMDCCXLIX"
	actual := intToRoman(num)
	assert.Equal(t, expected, actual)

	num = 0
	expected = ""
	actual = intToRoman(num)
	assert.Equal(t, expected, actual)

	num = 400
	expected = "CD"
	actual = intToRoman(num)
	assert.Equal(t, expected, actual)
}
