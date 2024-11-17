package leetcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
返回 合并后的字符串 。

示例：
输入：word1 = "abc", word2 = "pqr"
输出："apbqcr"
*/
func mergeAlternately(word1, word2 string) string {
	if word1 == "" {
		return word2
	} else if word2 == "" {
		return word1
	}

	var resp string
	for i := 0; i < len([]rune(word1)); i++ {
		if i > (len([]rune(word2)) - 1) { // 说明 word2 已经没有了.
			resp += string([]rune(word1)[i])
		} else {
			resp += string([]rune(word1)[i]) + string([]rune(word2)[i])
		}
	}

	if len([]rune(word2)) > len([]rune(word1)) { // 说明上面遍历的不够，将 word2 全部加进去.
		for i := len([]rune(word1)); i < len([]rune(word2)); i++ {
			resp += string([]rune(word2)[i])
		}
	}

	return resp
}

/*
对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。

示例：
输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
*/
func gcdOfStrings(str1 string, str2 string) string {
	/*
		结题思路：
		1、如果子串能被 str1 和 str2 除尽，则 len(str1)/len(subStr) = 0 && len(str2)/len(subStr) = 0.
		2、判断这个子串是不是能被除尽的子串，需要判断将 n 个 subStr 拼接起来后，是否等于原串.
	*/
	return ""
}

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

示例：
输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
*/
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	// 第一个不为空格的字符，为开始，再遍历到空格时，结束.
	words := make([]string, 0)
	var isStart bool
	var word string
	for _, val := range []rune(s) {
		if string(val) != " " {
			word = word + string(val)
			isStart = true
		} else {
			if isStart { // 已经开始遍历了，且遇到了空格，则结束.
				isStart = false
				words = append(words, word)
				word = ""
			}
		}
	}

	// 倒序遍历.
	var resp string
	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 {
			resp += words[i]
		} else {
			resp = resp + words[i] + " "
		}
	}
	return resp
}

/*
判断字符串是否是有效的 Ipv4 地址.
*/
func validIPAddress(queryIP string) bool {
	var subStr string
	for k, v := range queryIP {
		if string(v) == "." || k == (len(queryIP)-1) {
			if k == (len(queryIP) - 1) {
				subStr += string(v)
			}
			val, err := strconv.ParseInt(subStr, 10, 64)
			if err != nil {
				return false
			}

			if val < 0 || val > 255 {
				return false
			}
			subStr = ""
		} else {
			subStr += string(v)
		}
	}

	return true
}

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

示例：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for _, v := range nums2 {
		for i := 0; i < len(nums1); i++ {
			if v > nums1[i] { // 找到了应该插入的位置.
				nums1 = append(nums1[:i], nums1[i:]...)
				nums1[i] = v
			}
		}
	}

	fmt.Println(nums1)
}

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
*/
func rotate(nums []int, k int) {
	resp := make([]int, len(nums))
	for idx, v := range nums {
		var newIdx int
		if idx+k >= len(nums) {
			newIdx = (idx + k) % len(nums)
		} else {
			newIdx = idx + k
		}
		resp[newIdx] = v
	}

	fmt.Println(resp)
}

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。
*/
func romanToInt(s string) int {
	special := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	common := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int
	for k, v := range s {
		tmpStr := string(v)
		// 判断时，往下多判断一位，看是否是特殊字符.
		if k != len(s)-1 {
			tmpStr += string([]byte(s)[k+1])
		}

		if specialResult, ok := special[tmpStr]; ok {
			result += specialResult
			k += 2
		} else {
			result += common[tmpStr]
		}
	}

	return result
}

/*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
*/
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	isalnum := func(ch byte) bool {
		return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
	}

	var i int
	j := len(s) - 1
	for {
		if i == j { // 当 i 和 j 相遇，表示是回文串.
			return true
		}

		if i > j {
			return false
		}

		if !isalnum(s[i]) {
			i++
			continue
		} else if !isalnum(s[j]) {
			j--
			continue
		}

		// 判断两个字符是否相等，若不等，则不是回文串.
		if s[i] != s[j] {
			return false
		}
	}
}

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
*/
func minWindow(s string, t string) string {
	// isContain 定义一个 s1 是否 t 的方法.
	isContain := func(s1 string) bool {
		for _, tv := range t {
			if !strings.Contains(s, string(tv)) {
				return false
			}
		}

		return true
	}

	// 定义一个滑动窗口，找到窗口内包含字符串 t 的字符串.
	var i, j int
	var minStr string
	for ; j < len(s); j++ {
		if isContain(s[i:j]) { // 若包含了字符串 t，i 向右滑动，找到最小的字符串.
			if len(s[i:j]) < len(minStr) {
				minStr = s[i:j]
			}

			// i 向右滑动，找到最小的子串.
			i++
			for ; i < j; i++ {
				if isContain(s[i:j]) {
					if len(s[i:j]) < len(minStr) {
						minStr = s[i:j]
					}
				}
			}

		}
	}

	return minStr
}

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
示例：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	var dupCount int // 重复次数.
	for k, num := range nums {
		if k == 0 {
			continue
		}

		// 由于已经非严格递增，判断下一个数字，若不等于前面的数字，说明没有重复.
		if num == nums[k-1] {
			dupCount++
		}
	}

	return len(nums) - dupCount
}

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
示例：
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
*/
func removeDuplicates2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	var dupCount int // 重复次数.
	for k, num := range nums {
		if k <= 1 {
			continue
		}

		// 由于已经非严格递增，判断下一个数字，若不等于前面的数字，说明没有重复.
		if num == nums[k-1] && num == nums[k-2] { // 判断是否重复的条件需要改变.
			dupCount++
		}
	}

	return len(nums) - dupCount
}

/*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

示例：
输入：nums = [2,3,1,1,4]
输出：true
*/
func canJump(nums []int) bool {
	// 遍历数组中的每一个值，该值可达到的最远位置为：x + nums[x].
	// 遍历 nums[x] 前的每个值，其可到达的最远位置只要超过了数组最大值，即可以到达最后.
	for k, v := range nums {
		if k+v >= (len(nums) - 1) { // 说明当前值可以达到最后，返回 true.
			return true
		}

		// 否则，将数组减去 k 前面的值，判断是否能达到最后.
		return canJump(nums[k+1 : k+v+1])
	}

	return false
}

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
示例：
输入：s = "   fly me   to   the moon  "
输出：4
*/
func lengthOfLastWord(s string) int {
	var lastWord string
	var isStart bool
	var wordLen int
	for _, v := range s {
		if string(v) != " " {
			lastWord += string(v)
			if !isStart { // 开始扫描到第一个字符.
				isStart = true
			}
		} else { // 又扫描到了空格.
			if isStart { // 说明前面已经开始了，到这里结束.
				isStart = false
				wordLen = len(lastWord)
				lastWord = ""
			}
		}
	}

	return wordLen
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

输入：strs = ["flower","flow","flight"]
输出："fl"
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	commonPrefix := strs[0]
	findPrefixFn := func(word string) string {
		if commonPrefix == "" {
			return ""
		}

		var prefix string
		for k, v := range word {
			if string(v) != string(commonPrefix[k]) { // 说明已经不是公共前缀了.
				return prefix
			}

			prefix += string(v)
		}

		return prefix
	}

	for i := 1; i < len(strs); i++ {
		commonPrefix = findPrefixFn(strs[i])
	}
	return commonPrefix
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func merge2(intervals [][]int) [][]int {
	resp := make([][]int, 0)
	resp = append(resp, intervals[0])
	for k, v := range intervals {
		left := resp[len(resp)-1] // 将左边作为已合并的区间.
		if k == 0 {
			continue
		}

		if left[1] >= v[0] { // 有重叠区间，合并.
			resp[len(resp)-1] = []int{left[0], v[1]} // 合并.
		} else { // 没有重叠区间，直接 append.
			resp = append(resp, v)
		}
	}

	return resp
}

/*
给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为 更加简洁的规范路径。

在 Unix 风格的文件系统中规则如下：

一个点 '.' 表示当前目录本身。
此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
返回的 简化路径 必须遵循下述格式：

始终以斜杠 '/' 开头。
两个目录名之间必须只有一个斜杠 '/' 。
最后一个目录名（如果存在）不能 以 '/' 结尾。
此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
返回简化后得到的 规范路径 。

示例：
输入：path = "/home/user/Documents/../Pictures"
输出："/home/user/Pictures"
*/
func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	resp := make([]string, 0)
	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		if dir == ".." { // 回到父目录.
			resp = resp[:len(resp)-1]
		} else {
			resp = append(resp, dir)
		}
	}

	fmt.Println(resp)
	return "/" + strings.Join(resp, "/")
}

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
实现 MinStack 类:
MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。

示例：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]
*/

// MinStack 最小栈.
type MinStack struct {
	len     int
	minData int
	data    []int // 存储数据.
}

func Constructor(len int) MinStack {
	return MinStack{
		len:     len,
		minData: 0,
		data:    make([]int, 0, len),
	}
}

func (s *MinStack) push(val int) error {
	if s.len == len(s.data) {
		return errors.New("stask is full")
	}

	if val < s.minData {
		s.minData = val
	}

	// 入栈.
	s.data = append(s.data, val)
	s.len++
	return nil
}

func (s *MinStack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stask is empty")
	}

	resp := s.data[s.len-1]
	if resp == s.minData {
		// 遍历栈中的最小值，重新赋值 minData.
		var minData int
		for _, v := range s.data {
			if minData > v {
				minData = v
			}
		}
		s.minData = minData
	}
	return resp, nil
}

func (s *MinStack) getMin() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stask is empty")
	}

	return s.minData, nil
}

/*
给你一个链表的头节点 head ，判断链表中是否有环。
*/
func hasCycle(head *ListNode) bool {
	// 采用 map 存储当前节点被访问过.
	nodeMap := make(map[*ListNode]struct{})
	tmpNode := head
	for {
		if tmpNode.Next == nil {
			return false // 没有环.
		}

		_, ok := nodeMap[tmpNode]
		if ok {
			return true
		}

		nodeMap[tmpNode] = struct{}{}
		tmpNode = tmpNode.Next
	}
}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var respList *ListNode
	tmpL1 := l1
	tmpL2 := l2
	var carry int
	for tmpL1 != nil || tmpL2 != nil { // 没有达到尾节点，需要遍历.
		var sum int
		var l1Data, l2Data int
		if tmpL1 != nil {
			l1Data = tmpL1.Val
			tmpL1 = tmpL1.Next
		}

		if tmpL2 != nil {
			l2Data = tmpL2.Val
			tmpL2 = tmpL2.Next
		}
		sum = l1Data + l2Data + carry // l1+l2+carry 除了两个值外，还要加上进位值.
		carry = sum % 10              // 是否需要进位.

		// 放入 resp 链表.
		if respList == nil {
			respList = &ListNode{
				Val:  sum,
				Next: nil,
			}
		} else {
			tmpNode := &ListNode{
				Val:  sum,
				Next: nil,
			}
			respList.Next = tmpNode // 将新节点挂到原节点的 next 下.
			respList = tmpNode      // 指针滑动到最后一个节点.
		}
	}
	if carry > 0 { // 判断是否有进位值（加出来的大于 10 的进位）.
		respList.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
		respList = respList.Next
	}

	return respList
}

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
示例：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return nil
}
