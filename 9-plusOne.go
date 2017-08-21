package main

// 一个非负整数， 用一个数组来表示每一位上的数字，如99 表示为{9,9} 实现数字加一之后数字的数组表示
func plusOne(digits []int) []int {
	l := len(digits)
	need_add := true
	for i := l - 1; i >= 0; i-- {
		if need_add {
			digits[i] = digits[i] + 1
		}

		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			need_add = true
		} else {
			need_add = false
		}
	}
	if need_add {
		digits = append([]int{1}, digits...)
	}
	return digits
}
