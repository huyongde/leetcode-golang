package main

import "fmt"

// 学会换位思考， 第一个思路是把所有等于val的删除， 这个思路不好实现，但是反向思维之后，
// 只获取和val不相等的，来填充数组前cnt个元素的方案则比较好实现
func removeElement(nums []int, val int) int {
	fmt.Println(nums)
	l := len(nums)
	cnt := 0
	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}
