package leetcode

/*
给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列，
请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，
则 1 <= index1 < index2 <= numbers.length 。
以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	resp := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for h := j + 1; h < len(nums); h++ {
				if nums[i]+nums[j]+nums[h] == 0 {
					resp = append(resp, []int{nums[i], nums[j], nums[h]})
				}
			}
		}
	}

	return resp
}

// -------------几种排序算法-------------.

// bubbleSort 冒泡排序.
// 每两个比较一下，将最大的往后放.
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ { // 一共要做这么多次.
		for j := 0; j < len(nums)-i-1; j++ { // 每次移动完成后，后面的一个数字，就不用再移动了.
			if nums[j] > nums[j+1] { // 交换.
				tmp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
			}
		}
	}

	return nums
}

// selectSort 选择排序.
// 每次都把最小的放在前面  []int{5, 2, 10, 1, 3, 12}.
// 和冒泡排序不同的是，每次循环是找数组中最小值的下标，最后交换一次，交换到前面.
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ { // 一共要做这么多次.
		minIdx := i
		for j := i; j < len(nums)-1; j++ { // 每次比较时，将最小值放到 i 的位置上.
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		// 找到最小值的位置，做一次交换.
		tmp := nums[i]
		nums[i] = nums[minIdx]
		nums[minIdx] = tmp
	}

	return nums
}

// insertSort 插入排序.
// 维护一个有序区间，每次将一个值插入到有序区间内.
func insertSort(nums []int) []int {
	resp := make([]int, 0, len(nums))
	resp = append(resp, nums[0])
	for i := 1; i < len(nums); i++ {
		j := len(resp) - 1
		for ; j >= 0; j-- {
			if resp[j] > nums[i] { // 结果数组里的值，大于当前值，则往后移动.
				resp[j+1] = resp[j]
			} else { // 找到插入的位置.
				break
			}
		}

		resp[j] = resp[i]
	}
	return resp
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	// 选定一个值作为待比较值，right 从右往左右，left 从左往右走，left、right 相遇时，则找到了待比较值的位置.
	beComparedData := nums[left]
	for left < right {
		if nums[right] < beComparedData {
			tmp := nums[left]
			nums[left] = nums[right]
			nums[right] = tmp
			right--
		} else if nums[left] > beComparedData {
			tmp := nums[right]
			nums[right] = nums[left]
			nums[left] = tmp
			left++
		}
	}

	// 找到了当前值所在的位置.
	nums[left] = beComparedData

	// 递归遍历左边和右边.
	quickSort(nums[0:left], 0, left-1)
	quickSort(nums[left+1:], left+1, len(nums[left+1:]))
}
