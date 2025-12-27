package q3maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	max := 0;
	count := 0;
	for i := range nums {
		if nums[i] == 1 {
			count++
			if max < count {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
