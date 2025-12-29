package hindex

import (
	"slices"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
  slices.Reverse(citations)
	count := 0
	for i := 1; i <= len(citations);i++ {
		if citations[i-1] >= i {
			count++
		} else {
			break
		}
	}
	return count
}
