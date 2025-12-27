package q2shufflethearray

func shuffle(nums []int, n int) []int {
	m := make([]int,0,2*n)
	for i := 0; i < n; i++ {
		m = append(m, nums[i])
		m = append(m, nums[n+i])
	}
	return m
}
