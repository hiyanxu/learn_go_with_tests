package array_slice

/**
两种初始化数组的方式：
1、[5]{1, 2, 3, 4, 5}
2、[...]{1, 2, 3, 4, 5}
*/
func Sum(numbers [5]int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}
