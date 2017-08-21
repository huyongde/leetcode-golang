package main

//import (
//	"fmt"
//)

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	//	fmt.Println(res)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			//			fmt.Println(i, j)
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}

func getRow(rowIndex int) []int {
	res := make([]int, 0)
	for i := 0; i < rowIndex+1; i++ {
		res = append(res, make([]int, 1)...)
		for j := i; j >= 0; j-- {
			if j == 0 || j == i {
				res[j] = 1
			} else {
				res[j] = res[j] + res[j-1]
			}
		}
	}
	return res
}
