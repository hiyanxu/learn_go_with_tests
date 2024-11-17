package sort

type insertSort struct{}

func (s *insertSort) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		j := i - 1
		for ; j > 0; j-- { // 从后往前找.
			if num < arr[j] { // 当前数据小于已排序队列的最后一个数据，则交换位置，向后挪一位.
				arr[i] = arr[j]
			} else { // 小于该值，直接跳出循环，说明找到了对应的插入位置.
				break
			}
		}

		arr[j+1] = num
	}

	return arr
}
