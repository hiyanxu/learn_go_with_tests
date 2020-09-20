package array_slice

/*
slice添加元素的几种操作：
1、正常添加一个元素。
2、给一个切片追加一个切片（此时需要解包操作）。
3、在切片中间第i个元素后添加元素（两种方式：（1）需要创建临时切片。（2）利用copy和append，不创建临时切片。）
*/

// 正常添加一个元素操作
func AppendNormal(sliceP []int, i int) []int {
	return append(sliceP, i)
}

// 追加切片
func AppendSlice(sliceP []int, sliceA []int) []int {
	// 需要在sliceA后面增加...，进行解包
	return append(sliceP, sliceA...)
}

// 在切片中间插入一个元素  需要创建临时切片
// sliceP 被插入的切片
// index 插入位置
// val 待插入的值
func InsertOne1(sliceP []int, index int, val int) []int {
	// 利用append函数，先将index及index位置后的元素，追加到一个临时切片中，然后再将该临时切片追加到第index的位置
	// append([]int{val}, sliceP[index:]...)：首先创建一个临时[]int{val}切片，将val加进来，然后将sliceP切片index位置后的数据加到临时切片后
	// 将上面创建的切片加到sliceP index下标后面，完成操作
	return append(sliceP[:index], append([]int{val}, sliceP[index:]...)...)
}

// 在切片中间插入一个元素  无需创建临时切片
// sliceP 被插入的切片
// index 插入位置
// val 待插入的值
func IndexOne2(sliceP []int, index int, val int) []int {
	// 利用copy和append函数

	// 利用append进行扩容
	sliceP = append(sliceP, 0)

	// 利用copy函数，将index下标及最后所有的数据，复制到index+1下标（向右移动一个位置）
	copy(sliceP[index+1:], sliceP[index:])

	// 将val赋值给index.
	sliceP[index] = val

	return sliceP
}
