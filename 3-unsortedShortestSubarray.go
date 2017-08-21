package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	start := 0
	has_start := false
	has_end := false
	max := nums[0]
	max2 := nums[0]
	end := 0
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < nums[i-1] || (has_start == true && nums[i] == nums[i-1] && nums[i] != max) {

			if has_start == false {
				start = i - 1
				max2 = nums[i-1]
			}
			has_start = true
			end = i
			has_end = true
		} else if has_start == true {
			if nums[i] >= max2 {
				max2 = nums[i]
			} else {
				end = i
				has_end = true
			}
		}
		if has_start && has_end {
			for j := start - 1; j >= 0; j-- { // 检查start元素之前有没有比当前end元素大的，若有则需要移动start
				if nums[j] > nums[i] {
					start = j
				}
			}
		}

	}
	fmt.Println(start, end)
	if end == start && start == 0 {
		return 0
	}
	return end - start + 1
}
