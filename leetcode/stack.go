package leetcode

import (
	"errors"
)

// stack 数组栈.
type stack struct {
	len  int
	data []string
}

func newStack(len int) *stack {
	return &stack{
		len:  len,
		data: make([]string, 0, len),
	}
}

func (s *stack) push(str string) error {
	if len(s.data) > s.len {
		return errors.New("stask is full")
	}

	s.data = append(s.data, str)
	return nil
}

func (s *stack) pop() (string, error) {
	if len(s.data) == 0 {
		return "", errors.New("stack is empty")
	}

	resp := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return resp, nil
}

/*
给你一个包含若干星号 * 的字符串 s 。
在一步操作中，你可以：
选中 s 中的一个星号。
移除星号 左侧 最近的那个 非星号 字符，并移除该星号自身。
返回移除 所有 星号之后的字符串。

示例：
输入：s = "leet**cod*e"
输出："lecoe"
解释：从左到右执行移除操作：
- 距离第 1 个星号最近的字符是 "leet**cod*e" 中的 't' ，s 变为 "lee*cod*e" 。
- 距离第 2 个星号最近的字符是 "lee*cod*e" 中的 'e' ，s 变为 "lecod*e" 。
- 距离第 3 个星号最近的字符是 "lecod*e" 中的 'd' ，s 变为 "lecoe" 。
不存在其他星号，返回 "lecoe" 。
*/
func removeStars(s string) string {
	if s == "" {
		return ""
	}

	st := []rune{}
	for _, c := range s {
		if c == '*' { // 若遇到了 * 号，则将已入数组的字符删除最后一个.
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

示例：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

输入：s = "3[a2[c]]"
输出："accaccacc"
*/
func decodeString(s string) string {
	if s == "" {
		return ""
	}

	type RepeatLetter struct {
		letter string
		count int
	}

	// 字符串入栈.
	stc := newStack(len(s))
	for _, v := range s {
		_ = stc.push(string(v))
	}

	// 示例："3[a2[c]]"
	// 挨个出栈，凑成 [c] 格式时，对其按照前面的数字重复一定的次数.
	repeatLetter := make([]*RepeatLetter, 0)
	var isFindRight bool
	for {
		letter, err := stc.pop()
		if err != nil {
			break
		}

		if letter == "]" {
			isFindRight = true
		} else {
			if isFindRight && letter != "[" {  // 找到了右边的符号，且当前不是[，则认为找到了待重复的字符.
				repeatLetter = append(repeatLetter, &RepeatLetter{
					letter: letter,
				})
			} else
		}
	}
}
