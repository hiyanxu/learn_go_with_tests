package alg

import (
	"errors"
	"fmt"
)

/*
一个二维数组，二维数组中的每个数都是非负数，要求从左上角走到右下角，每一步只能向右或者向下。沿途经过的数字要累加起来。返回最小的路径和。
例
1 3 5 9
8 1 3 4
5 0 6 1
8 8 4 0

路径1,3,1,0,6,1,0就是最小路径和，返回12

1 3 500 9
8 100 3 4
5 0 6 1
8 8 4 0
*/
func handle(arr [][]int) int {
	var index int
	var index2 int
	var sum int
	for i := 0; i < len(arr); i++ {
		if arr[index][index2] < arr[index][index2+1] {
			sum += arr[index][index2]
			index2 += 1
		} else {
			sum += arr[index][index2+1]
			index += 1
		}
	}

	return sum
}

/*
*
给一个字符串 str ，包含多个单词， 单词之间以及字符串前后可能有多个#。要求：

	1	将字符串中的单词提取出来，按输入串中出现的次序，做一个逆序排列， 单词之间用空格分割
	2	请不要使用库函数来分割字符串

例如， 输入：str = "##Please##show####me#the####code" 输出：str = "code the me show Please"
*/
type node struct {
	word string
	next *node
}

func (n *node) push(word string) {
	newNode := &node{
		word: word,
		next: nil,
	}
	n.next = newNode
}

func (n *node) rangeStack() {
	for {
		if n.next == nil {
			break
		}
	}
}

func (n *node) pop() string {
	var preNode *node
	for {
		if n.next == nil {
			preNode = n
			preNode.next = nil
			return n.word
		}
	}
}

func handleStr(str string, sep string) {
	words := make([]string, 0)
	var word string
	var isStart bool // 标记是否开始扫描到了单词.
	for key, val := range []rune(str) {
		if string(val) != sep {
			isStart = true
			word += string(val)

			if key == len([]rune(str))-1 {
				words = append(words, word)
			}
		} else {
			if isStart {
				words = append(words, word)
				isStart = false
				word = ""
			}
		}
	}

	fmt.Println(words)
}

/*
1->2->3->1
*/

type ListNode struct {
	data int
	next *ListNode
}

func handleList(l *ListNode) {
}

type ParentInfo struct {
	ID       int
	ParentID int
	Val      string
	Childs   []int
}

func GetParentInfo(id int) *ParentInfo {
	return &ParentInfo{
		ID:       id,
		ParentID: 0,
		Val:      "",
		Childs:   getChild(id),
	}
}

func getChild(id int) []int {
	child := getChild(id)
	if len(child) == 0 {
		return []int{}
	}

	childs := make([]int, 0)
	for _, childID := range child {
		childs = append(childs, getChild(childID)...)
	}
	return childs
}

/*
两个有序数组的排序
[1,3,6,7,8]
[2,5,9]
*/
func sort(arr1, arr2 []int) []int {
	resp := make([]int, 0, len(arr2)+len(arr1))
	var i, j int
	for {
		// 当任意一个下标达到最后，直接退出.
		if i == len(arr1) || j == len(arr2) {
			break
		}

		if arr1[i] <= arr2[j] { // 小的数据往里追加.
			resp = append(resp, arr1[i])
			i++
		} else {
			resp = append(resp, arr2[j])
			j++
		}
	}

	// 将剩余的数组数据追加到 resp 中.
	if i <= len(arr1)-1 {
		for ; i < len(arr1); i++ {
			resp = append(resp, arr1[i])
		}
	}
	if j <= len(arr2)-1 {
		for ; j < len(arr2); j++ {
			resp = append(resp, arr2[j])
		}
	}

	return resp
}

type LinkNode struct {
	data int
	next *LinkNode
}

type lruCache struct {
	data   map[string]string
	keyIdx []string
	cap    int
}

func NewLruCache(cap int) *lruCache {
	return &lruCache{
		data:   make(map[string]string, cap),
		keyIdx: make([]string, 0, cap),
		cap:    cap,
	}
}

func (c *lruCache) Get(key string) (string, error) {
	val, ok := c.data[key]
	if !ok {
		return "", errors.New("key not found")
	}

	// 将 key 放到最后面.
	var keyIdx int
	for idx, tmpKey := range c.keyIdx {
		if tmpKey == key {
			keyIdx = idx
		}
	}

	c.keyIdx = append(c.keyIdx[:keyIdx], c.keyIdx[keyIdx:]...)
	c.keyIdx = append(c.keyIdx, key)
	return val, nil
}

func (c *lruCache) Set(key string, val string) error {
	if c.cap == len(c.keyIdx) { // 缓存已满，淘汰最久未使用的数据.
		err := c.Delete()
		if err != nil {
			return err
		}
	}

	c.data[key] = val
	c.keyIdx = append(c.keyIdx, key)
	return nil
}

func (c *lruCache) Delete() error {
	if len(c.keyIdx) == 0 {
		return errors.New("cache is empty")
	}

	key := c.keyIdx[0]
	delete(c.data, key)
	c.keyIdx = c.keyIdx[1:]
	return nil
}

/*
二维数组，值横向递增，纵向递增，查询某个值.
1 3 5 8
2
4
*/
func searchData(arr [][]int, targetVal int) []int {
	var i, j int
	for {
		if i > len(arr) || j > len(arr[0]) {
			break
		}

		for ; j < len(arr[i]); j++ {
			if arr[i][j] == targetVal {
				return []int{i, j}
			}

			if arr[i][j] > targetVal { // 说明在横向后面, j++ 往后走.  // 说明在纵向， i++ 往下走.
				i = i + 1
			}
		}
	}

	return nil
}
