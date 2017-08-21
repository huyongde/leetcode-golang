package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	nums := flowerbed
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			continue
		}
		if (i-1 < 0 || (i-1 >= 0 && nums[i-1] == 0)) && (i+1 >= l || (i+1 < l && nums[i+1] == 0)) {
			nums[i] = 1
			n--
		}
		// fmt.Println(i, n)
		if n <= 0 {
			return true
		}

	}
	return false
}
