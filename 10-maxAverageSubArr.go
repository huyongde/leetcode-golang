package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	l := len(nums)
	max := 0
	if l <= k {
		for i := 0; i < l; i++ {
			max += nums[i]
		}
		return float64(max) / float64(k)
	}
	fmt.Println(1111)
	tmp := 0
	for j := 0; j < k; j++ {
		max += nums[j]
	}
	for i := 1; i <= l-k; i++ {
		tmp = 0
		for j := i; j < i+k; j++ {
			tmp += nums[j]
		}
		fmt.Println(max, tmp)
		if tmp > max {
			max = tmp
		}
	}
	return float64(max) / float64(k)
}
