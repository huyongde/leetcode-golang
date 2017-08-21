package main

import "fmt"

// K-diff Pairs in Array

//Given an array of integers and an integer k, you need to find the number of unique k-diff pairs in the array. Here a k-diff pair is defined as an integer pair (i, j), where i and j are both numbers in the array and their absolute difference is k.

//Example 1:
//Input: [3, 1, 4, 1, 5], k = 2
//Output: 2
//Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
//Although we have two 1s in the input, we should only return the number of unique pairs.
//Example 2:
//Input:[1, 2, 3, 4, 5], k = 1
//Output: 4
//Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
//Example 3:
//Input: [1, 3, 1, 5, 4], k = 0
//Output: 1
//Explanation: There is one 0-diff pair in the array, (1, 1).
//Note:
//The pairs (i, j) and (j, i) count as the same pair.
//The length of the array won't exceed 10,000.
//All the integers in the given input belong to the range: [-1e7, 1e7].
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	l := len(nums)
	if l == 1 {
		return 0
	}
	m1 := make(map[int]int, l) // map来存储数字出现的次数
	for i := 0; i < l; i++ {
		if _, ok := m1[nums[i]]; ok {
			m1[nums[i]]++
		} else {
			m1[nums[i]] = 1
		}
	}
	cnt := 0

	if k == 0 { // 当k=0时，只需统计出现次数大于等于2的数字，
		for _, v := range m1 {
			if v > 1 {
				cnt++
			}
		}
		return cnt
	}

	tmp1 := 0
	tmp2 := 0
	need_delete := false
	for key, _ := range m1 {
		need_delete = false
		tmp1 = key - k
		tmp2 = key + k
		if _, ok := m1[tmp1]; ok {

			cnt++
			need_delete = true
		}
		if _, ok := m1[tmp2]; ok {
			cnt++
			need_delete = true
		}
		if need_delete == true { // 一个数字访问之后，后续就不能在访问， 所以要从map中删除掉。delete(m, key)
			fmt.Println(11111)
			delete(m1, key)
		}

	}
	return cnt
}
