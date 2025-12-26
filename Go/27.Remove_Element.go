package main

func removeElement(nums []int, val int) int {
    a := 0
    for i, num := range nums {
			if val != nums[i] {
				nums[a] = num
				a++
			}
		}
		return a

		// nums = slices.DeleteFunc(nums, func(i int) bool {
		// 	return i == val
		// })

		// return len(nums)
}
