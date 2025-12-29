package hindex

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	fmt.Println(citations)
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
