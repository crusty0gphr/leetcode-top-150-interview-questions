package letter_combination_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombination(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	actual := letterCombinations(digits)
	assert.Equal(t, expected, actual)

	digits = ""
	expected = nil
	actual = letterCombinations(digits)
	assert.Equal(t, expected, actual)

	digits = "2"
	expected = []string{"a", "b", "c"}
	actual = letterCombinations(digits)
	assert.Equal(t, expected, actual)
}
