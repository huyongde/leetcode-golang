package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1}
	nums = []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 7}
	//fmt.Println(removeDuplicates(nums))
	nums = []int{1, 1, 2, 2}
	fmt.Println(searchInsert(nums, 1))
	fmt.Println(nums)

}
