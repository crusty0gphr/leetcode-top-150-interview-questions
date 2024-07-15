package text_justification

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var line []string
	lineWidth := 0
	i := 0

	for i < len(words) {
		word := words[i]
		if lineWidth+len(word)+len(line) > maxWidth {
			extraSpaces := maxWidth - lineWidth

			spacesPerWord := extraSpaces / max(1, len(line)-1) // avoid 0 division
			reminder := extraSpaces % max(1, len(line)-1)      // avoid 0 division

			for j := 0; j < max(1, len(line)-1); j++ { // avoid 0 loops
				line[j] = fmt.Sprintf("%s%s", line[j], strings.Repeat(" ", spacesPerWord))
				if reminder > 0 {
					line[j] = fmt.Sprintf("%s%s", line[j], " ")
					reminder--
				}
			}

			res = append(res, strings.Join(line, ""))
			line = []string{}
			lineWidth = 0
		}

		line = append(line, word)
		lineWidth += len(word)
		i++
	}

	lastLine := strings.Join(line, " ")
	trailingSpaces := maxWidth - len(lastLine)
	lastLine = fmt.Sprintf("%s%s", lastLine, strings.Repeat(" ", trailingSpaces))
	res = append(res, lastLine)

	return res
}
