package h_index

import (
	"slices"
)

func hIndex(citations []int) int {
	// return hIndexWithSorting(citations)
	return hIndexWithMap(citations)
}

func hIndexWithMap(citations []int) int {
	maxHIndex := len(citations)
	cMap := make(map[int]int)

	for _, citation := range citations {
		if citation >= maxHIndex {
			cMap[maxHIndex]++
		} else {
			cMap[citation]++
		}
	}

	highest := maxHIndex
	count := 0
	for i := maxHIndex; i > 0; i-- {
		count += cMap[i]
		if count < highest {
			highest--
		}
	}

	return highest
}

func hIndexWithSorting(citations []int) int {
	slices.Sort(citations) // O(n)

	n := len(citations)
	for i, citation := range citations {
		hIdx := n - i
		if hIdx <= citation {
			return n - i
		}
	}
	return 0
}
