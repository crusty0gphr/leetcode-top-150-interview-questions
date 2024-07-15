package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	input := "XXVII"
	expected := 27

	res := romanToInt(input)
	assert.Equal(t, expected, res)

	input = "III"
	expected = 3

	res = romanToInt(input)
	assert.Equal(t, expected, res)

	input = "LVIII"
	expected = 58

	res = romanToInt(input)
	assert.Equal(t, expected, res)

	input = "MCMXCIV"
	expected = 1994

	res = romanToInt(input)
	assert.Equal(t, expected, res)
}
