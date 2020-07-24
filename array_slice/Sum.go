package array_slice

/**
两种初始化数组的方式：
1、[5]{1, 2, 3, 4, 5}
2、[...]{1, 2, 3, 4, 5}

for循环遍历数组：
	通过传统for循环可以遍历数组
range遍历数组：
	range会返回迭代数组，且返回数组元素的索引和值：key, value := range array

数组类型：
	数组的长度也是类型的一部分，不同长度的数组之间不可进行赋值

重构：
1、通过range遍历数组更加简洁。
*/
func Sum(numbers [5]int) int {
	sum := 0

	// 采用for循环
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	// 重构1：采用range array
	for _, num := range numbers {
		sum += num
	}

	return sum
}

/**

 */
func SumBySlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

/**
make([]int, len)：
	（1）切片需要通过make函数进行初始化后才可以使用；
	（2）使用append()动态增加，不通过sum[i]的方式
	未初始化直接使用会报错：panic: runtime error: index out of range [0] with length 0
*/
func SumAll(numbersToSum ...[]int) (sum []int) {
	len := len(numbersToSum)
	sum = make([]int, len)

	for i := 0; i < len; i++ {
		//sumOfNum := SumBySlice(numbersToSum[i])
		//sum = append(sum, sumOfNum)
		sum[i] = SumBySlice(numbersToSum[i])
	}

	return
}

/**
可以使用slice[low:high]获取部分切片
（1）包括左边不包括右边（若右边没有具体指定哪里结束，则包括最后） 长度：high - low  容量：底层数组长度 - low
*/
func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	length := len(numbersToSum)

	for i := 0; i < length; i++ {
		if len(numbersToSum[i]) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbersToSum[i][1:]
			sum = append(sum, SumBySlice(tail))
		}
	}

	return sum
}
