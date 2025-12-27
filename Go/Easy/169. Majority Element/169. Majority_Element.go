package main

import "sort"

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	sort.Ints(nums)
	for _, v := range nums {
		counts[v]++
	}
	max := -1
	idx := 0
	for k, v := range counts {
		if max < v {
			max = v
			idx = k
		}
	}
	return idx
}
