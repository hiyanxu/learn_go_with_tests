package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeAlternately(T *testing.T) {
	resp := mergeAlternately("abc", "pqrse")
	fmt.Println(resp)
}

func TestReverseWords(T *testing.T) {
	resp := reverseWords("  hello world  ")
	fmt.Println(resp)
}

func TestValidIPAddress(T *testing.T) {
	fmt.Println(validIPAddress("192.168.1.1"))
	fmt.Println(validIPAddress("192.168.1.266"))
}

func TestRotate(T *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func TestRemoveDuplicates2(t *testing.T) {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func TestMerge2(t *testing.T) {
	fmt.Println(merge2([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge2([][]int{{1, 4}, {4, 5}}))
}

func TestSimplifyPath(t *testing.T) {
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
	fmt.Println(simplifyPath("/home2//user2/Documents2/../Pictures2"))
}

func TestRomanToInt(t *testing.T) {
	fmt.Println(romanToInt("LVIII"))
}
