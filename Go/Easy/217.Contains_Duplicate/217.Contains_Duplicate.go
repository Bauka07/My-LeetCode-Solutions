package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
    for i:=0; i < len(nums)-1;i++ {
        if nums[i] == nums[i + 1] {
            return true
        }
    }
    sort.Ints(nums)
    for i:=0; i < len(nums)-1;i++ {
        if nums[i] == nums[i + 1] {
            return true
        }
    }
    return false
}
