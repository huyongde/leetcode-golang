package main

import (
	"fmt"
)

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

//For example,
//Given the following matrix:

//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//You should return [1,2,3,6,9,8,7,4,5].

// 分析题目， 只需要，按照， 向右， 向下， 向左， 向下 四个步骤，迭代执行就行。
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	res := []int{}
	is, ie, js, je := 0, m-1, 0, n-1
	for is <= ie && js <= je {
		// 向右遍历
		fmt.Println(is, ie, js, je)
		for j := js; j <= je; j++ {
			res = append(res, matrix[is][j])
		}
		is++
		// 向下遍历
		for i := is; i <= ie; i++ {
			res = append(res, matrix[i][je])
		}
		je--

		if is <= ie { // 只有在行还未全走一遍的时候去向左遍历
			// 向左遍历
			for j := je; j >= js; j-- {
				res = append(res, matrix[ie][j])
			}
			ie--
		}

		// 向上遍历
		if js <= je { // 只有在列还未全走完时，才去向上遍历
			for i := ie; i >= is; i-- {
				res = append(res, matrix[i][js])
			}
			js++
		}
		//		fmt.Println(res)

	}
	//	fmt.Println(len(res))
	return res
}
