package main

func twoSum(nums []int, target int) []int {
		result := make(map[int]int)
    for i, num := range nums {
        current := target - num
				if j, exists := result[current]; exists {
					return []int{j, i}
				}

				result[num] = i
    }

		return nil
}

