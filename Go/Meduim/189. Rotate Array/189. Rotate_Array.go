package main

func rotate(nums []int, k int) {
    if len(nums) < 2 {
        return
    }
    if k == 0 {
        return
    }

    k %= len(nums)

    slice := []int{}
    slice = append(slice, nums[len(nums) - k:]...)
    slice = append(slice, nums[:len(nums) -k]...)
    copy(nums, slice)
}
