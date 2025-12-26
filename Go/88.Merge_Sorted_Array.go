package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	combined := []int{}
	for i := 0; i < m; i++ {
		combined = append(combined, nums1[i])
	}
	for i := 0; i < n; i++ {
		combined = append(combined, nums2[i])
	}
	sort.Ints(combined)
	for i := 0; i < n+m; i++{
		nums1[i] = combined[i]
	}
}
