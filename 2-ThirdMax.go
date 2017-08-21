package main

import "fmt"

//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

//Example 1:
//Input: [3, 2, 1]

//Output: 1

//Explanation: The third maximum is 1.
//Example 2:
//Input: [1, 2]

//Output: 2

//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//Example 3:
//Input: [2, 2, 3, 1]

//Output: 1

//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.

func thirdMax(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n < 3 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	max := 0
	second := max
	third := max
	has_max, has_second, has_third := false, false, false //

	for i := 0; i < n; i++ {
		if nums[i] > max || has_max == false {
			//			if nums[i] == max {
			//				continue
			//			}
			fmt.Println("test1")
			third = second
			second = max
			max = nums[i]
			if has_second == true {
				has_third = true
			}
			if has_max == true {
				has_second = true
			}
			has_max = true

		} else if (nums[i] > second || has_second == false) && nums[i] != max {
			fmt.Println("test2")
			fmt.Println(has_second)
			//			if nums[i] == second {
			//				continue
			//			}
			third = second
			second = nums[i]
			if has_second == true {
				has_third = true
			}
			has_second = true
		} else if (nums[i] > third || has_third == false) && nums[i] != second && nums[i] != max {
			fmt.Println("test3")
			third = nums[i]
			has_third = true
		}
		fmt.Println(i, nums[i], max, second, third)
	}
	if third == second || third == max || second == max || has_third == false {
		return max
	}

	return third
}
