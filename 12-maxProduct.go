package main

import (
	"fmt"
)

// 连续的子数组乘积最大
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	max := nums[0]
	min := max
	res := max // 最大值和最小值以及最终结果都初始化为数组的第一个元素， 记录最小值，是因为可能负负得正然后取的一个最大值。
	for i := 1; i < n; i++ {
		tmp := min * nums[i] // 存储可能负负得正的结果，来和max比较

		min = Min(nums[i], min*nums[i])
		min = Min(min, max*nums[i]) // 更新以index为i的元素结尾的最小乘积

		max = Max(nums[i], max*nums[i])
		max = Max(max, tmp) // 更新以index为i的元素结尾的最大乘积

		res = Max(res, max) // 更新最新结果
		fmt.Println(i, min, max, res)
	}
	return res
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
