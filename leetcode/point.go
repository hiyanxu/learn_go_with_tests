package leetcode

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例：
输入: nums = [2,0,0,1,3,12]
输出: [2,1,3,12,0,0]
*/
func moveZeroes(nums []int) {
	right, left, n := 0, 0, len(nums)
	for right < n { // 右下标是否指向了最后.
		if nums[right] != 0 { // 找到了不为 0 的数据，做交换.
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例：
输入：s = "abc", t = "ahbgdc"
输出：true
*/
func isSubsequence(s string, t string) bool {
	if s == "" || t == "" {
		return false
	}

	var findedStr string
	var j int
	for i := 0; i < len([]rune(s)); i++ {
		for ; j < len([]rune(t)); j++ {
			if []rune(t)[j] == []rune(s)[i] { // 找到了某个字符.
				findedStr += string([]rune(s)[i])
				break // 跳出本层循环，下次时，j 直接从当前位置开始.
			}
		}
	}

	return findedStr == s
}
