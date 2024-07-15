package zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	expected := "PAHNAPLSIIGYIR"
	actual := convert(s, numRows)
	assert.Equal(t, expected, actual)

	s = "PAYPALISHIRING"
	numRows = 4
	expected = "PINALSIGYAHRPI"
	actual = convert(s, numRows)
	assert.Equal(t, expected, actual)

	s = "A"
	numRows = 1
	expected = "A"
	actual = convert(s, numRows)
	assert.Equal(t, expected, actual)
}
