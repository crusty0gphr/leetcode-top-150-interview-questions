package zigzag_conversion

import "strings"

func convert(str string, numRows int) string {
	if numRows == 1 || numRows >= len(str) {
		return str
	}

	rows := make([]strings.Builder, min(numRows, len(str)))
	currentRow := 0
	goingDown := false

	for _, char := range str {
		rows[currentRow].WriteRune(char)

		if currentRow == 0 {
			goingDown = true
		}

		if currentRow == numRows-1 {
			goingDown = false
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}
	return res.String()
}
