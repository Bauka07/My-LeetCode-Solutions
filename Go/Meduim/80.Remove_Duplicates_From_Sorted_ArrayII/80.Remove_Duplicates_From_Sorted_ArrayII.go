package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
			return len(nums)
	}
	a := 2
	for i:=2; i < len(nums); i++{
		if nums[i] != nums[a-2] {
			nums[a] = nums[i]
			a++
		}
	}
	return a
}
