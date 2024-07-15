package roman_to_integer

func romanToInt(number string) int {
	rtoi := 0
	rtoiLookup := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(number) - 2; i >= 0; i-- {
		currentChar := number[i]
		nextChar := number[i+1]

		currentVal := rtoiLookup[currentChar]
		nextVal := rtoiLookup[nextChar]

		if nextVal > currentVal {
			rtoi -= currentVal
		} else {
			rtoi += currentVal
		}
	}

	lastChar := number[len(number)-1]
	lastVal := rtoiLookup[lastChar]
	rtoi += lastVal

	return rtoi
}
