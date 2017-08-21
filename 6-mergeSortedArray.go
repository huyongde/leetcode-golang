package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = nums2
		return
	}
	if n == 0 {
		return
	}
	nums1 = append(nums1, nums2...)
	fmt.Println(nums1)
	j, k := m-1, n-1
	for i := m + n - 1; i >= 0 && k >= 0 && j >= 0; i-- {
		fmt.Println(i, j, k)
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
	if k > 0 {
		for i := k; i >= 0; i-- {
			nums1[k] = nums2[k]
		}
	}
	fmt.Println(nums1)
}
