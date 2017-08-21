package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums) // 先对数组进行排序
	fmt.Println(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ { // 遍历数组中每个元素
		if i > 0 && nums[i] == nums[i-1] { // 同一个元素处理一次
			continue
		}
		l, r := i+1, n-1
		for l < r { //
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			} else {
				tmp := []int{nums[i], nums[l], nums[r]}
				res = append(res, tmp)
				for nums[r-1] == nums[r] && l < r { // 同一个元素处理一次
					r--
				}
				for nums[l+1] == nums[l] && l < r { // 同一个元素处理一次
					l++
				}
				l++
				r--
			}
		}
	}
	return res
}
