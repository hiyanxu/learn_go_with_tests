package search

/*
递归实现二分查找.
*/
func search(arr []int, targetVal int, left, right int) int {
	i := (left + right) / 2
	if arr[i] == targetVal { // 判断当前值是否是目标值.
		return i
	}

	if arr[i] < targetVal { // 判断在队列的哪边.
		search(arr, targetVal, 0, i)
	} else {
		search(arr, targetVal, i, len(arr))
	}
	return -1
}
