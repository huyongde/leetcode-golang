package main

// 求和最大的连续子数组， 方案： 维护两个变量， 包括当前元素的最大值max, 在所有的max中选择出全局最大res
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	max := nums[0]
	res := max

	for i := 1; i < l; i++ {
		if nums[i] > max+nums[i] {
			max = nums[i]
		} else {
			max = max + nums[i]
		}
		if res < max {
			res = max
		}
	}
	return res
}
