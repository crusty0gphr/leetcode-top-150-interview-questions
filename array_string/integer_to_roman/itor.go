package integer_to_roman

import "strings"

func intToRoman(num int) string {
	var res strings.Builder
	rtoiLookup := []struct {
		roman   string
		integer int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	for _, lookup := range rtoiLookup {
		if num/lookup.integer > 0 {
			multiplier := num / lookup.integer
			res.WriteString(
				strings.Repeat(lookup.roman, multiplier),
			)
			num = num % lookup.integer
		}

	}

	return res.String()
}
