package array_slice

/**
切片删除：删除一个元素或删除一个切片
1、在尾部删除。
2、在开头删除。
3、在中间删除。
*/

// 在尾部删除
func DeleteOneAtTail(s []int) []int {
	return s[:len(s)-1]
}

// 在开头删除
func DeleteOneInBeginning(s []int) []int {
	/**
	两种方式：
	1、直接返回子切片，从原切片的下标1开始。
	2、利用append，将1之后的所有元素赋值给一个新的切片。
	*/
	//return s[1:]

	return append(s[:0], s[1:]...)
}

func DeleteOneInMiddle(s []int, index int) []int {
	// append：将index+1后的元素追加到index的位置
	return append(s[:index], s[index+1:]...)
}
