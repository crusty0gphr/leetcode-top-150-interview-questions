package letter_combination_of_a_phone_number

import "fmt"

var diToCh = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var res []string

	if len(digits) > 0 {
		return backtracking(digits, 0, "", res)
	}

	return res
}

func backtracking(
	digits string,
	numIdx int,
	currStr string,
	res []string,
) []string {
	if len(currStr) == len(digits) {
		res = append(res, currStr)
	} else {
		charSet := diToCh[digits[numIdx]]

		for _, char := range charSet {
			res = backtracking(
				digits,
				numIdx+1,
				fmt.Sprintf("%s%s", currStr, string(char)),
				res,
			)
		}
	}
	return res
}
