package main

import (
	"fmt"
)

//import "fmt"

//Given a sorted array, remove the duplicates in place
// such that each element appear only once and return the new length.

//Do not allocate extra space for another array,
//you must do this in place with constant memory.

//For example,
//Given input array nums = [1,1,2],

//Your function should return length = 2,
//with the first two elements of nums being 1 and 2 respectively.
//It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			for j := i; j < n-1; j++ {
				nums[j] = nums[j+1]
			}
			n = n - 1 // 每次移动元素，就意味着有重复的存在
			i = i - 1 // 移动完了还是要比较下当前元素和上一个是否相等
		}
		fmt.Println(i, nums[i], n)
	}
	return n
}
