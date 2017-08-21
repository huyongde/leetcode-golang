package main

import "fmt"

//func main() {
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	//	a = []int{1, 2}
//	k := 3
//	fmt.Println(a)
//	rotate(a, k)
//	fmt.Println(a)

//}

//把数组向右旋转移动K个元素， 比如：[1,2,3,4] 旋转移动2个元素， 变成[3,4,1,2], 数组长度是k的倍数时，外层for循环需要进行多次， 不是倍数时，外层for循环只需要一次
// Rotate an array of n elements to the right by k steps.

//For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

//Note:
//Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.

//[show hint]

//Related problem: Reverse Words in a String II
func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	var i, j = 0, 0
	var tmp, tmp1 int
	var n = 0
	//	if l > k {
	for m := 0; m < l; m++ {
		fmt.Println("m:", m)
		i = m
		for {

			n = n + 1
			if n > l {
				break
			}
			j = (i + k) % l
			if i == m {
				tmp = nums[j]
				nums[j] = nums[i]
			} else {
				tmp1 = nums[j]
				nums[j] = tmp
				tmp = tmp1
			}
			i = j
			if j == m {
				break
			}
		}
		if n > l {
			break
		}
	}

	//	}
}
