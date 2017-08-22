package main

//import "fmt"
// 要考虑到两个边界情况，数组为空， 或者target比数组所有的元素都大
func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return l
}
