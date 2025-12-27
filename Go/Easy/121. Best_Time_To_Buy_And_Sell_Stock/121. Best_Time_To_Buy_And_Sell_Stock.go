package main

func maxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if maxProfit < prices[i] - min  {
			maxProfit = prices[i] - min
		}
	}
	return maxProfit
}
