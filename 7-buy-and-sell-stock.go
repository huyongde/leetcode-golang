package main

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	max := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
