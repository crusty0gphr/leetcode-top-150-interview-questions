package longest_common_prefix

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix strings.Builder
	sort.Strings(strs)

	first, last := strs[0], strs[len(strs)-1]
	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] != last[i] {
			return prefix.String()
		}
		prefix.WriteByte(first[i])
	}

	return prefix.String()
}
