package alg

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

var keyNotFound = errors.New("key not found")

type listNode struct {
	key  string
	pre  *listNode
	next *listNode
}

type cache struct {
	cap  int
	data map[string]string
	head *listNode
}

func NewCache(cap int) *cache {
	return &cache{
		cap:  cap,
		data: make(map[string]string, cap),
		head: &listNode{
			pre:  nil,
			next: nil,
		},
	}
}

func (c *cache) Get(key string) (string, error) {
	// 判断是否为空.
	if len(c.data) == 0 {
		return "", errors.New("cache is empty")
	}

	// 查找某个节点.
	val, ok := c.data[key]
	if !ok {
		return "", keyNotFound
	}

	// 从链表中找到该节点，做链表修改.
	var preNode *listNode
	nextNode := c.head.next
	currentNode := c.head
	for {
		if currentNode.key == key { // 找到了该节点.
			// 做一个删除操作，直接让 pre 节点指向 next 节点.
			preNode.next = nextNode
			break
		}

		// 往下找相关节点.
		preNode = currentNode          // 前一个节点为当前节点.
		currentNode = currentNode.next // 指向下一个节点.
		nextNode = currentNode.next    // 下一个节点的下一个节点.
	}

	// 把当前这个节点，放到最后.
	lastNode := c.head
	for {
		if lastNode.next == nil { // 找到了最后的节点.
			break
		}

		lastNode = lastNode.next
	}
	lastNode.next = currentNode
	currentNode.next = nil

	return val, nil
}

// getCommonStr 获取两个字符串的最长公共子串.
func getCommonStr(str1, str2 string) (int, string) {
	var lognSubStr string
	var resp string
	var i int
	var commonStrLen int
	for {
		if i >= len([]byte(str1)) || i >= len([]byte(str2)) {
			break
		}

		if []byte(str1)[i] == []byte(str2)[i] { // 两个字符相等时，说明找到了公共子串.
			lognSubStr += string([]byte(str1)[i])
		} else {
			if commonStrLen < len(lognSubStr) {
				commonStrLen = len(lognSubStr)
				resp = lognSubStr
			}
			lognSubStr = ""
		}
		i++
	}

	return commonStrLen, resp
}

func GetSumValFromFile(filePath string) int {
	fileReader, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("read file err: %+v", err)
	}
	defer fileReader.Close()

	scaner := bufio.NewScanner(fileReader)
	var sumVal int
	for scaner.Scan() {
		nums := getNumsByScaner(scaner.Text())
		sumVal += handleNums(nums)
	}

	return sumVal
}

func getNumsByScaner(str string) []int {
	strNums := strings.Split(str, ",")
	nums := make([]int, 0, len(strNums))
	for _, strNum := range strNums {
		num, err := strconv.ParseInt(strNum, 10, 64)
		if err != nil {
			fmt.Printf("parse string: %s num err: %+v", strNum, err)
			continue
		}

		nums = append(nums, int(num))
	}

	return nums
}

func handleNums(nums []int) int {
	ch := make(chan int, len(nums))
	var sumVal int64
	done := make(chan struct{})
	// go 出去一个异步的 goroutine，用于处理最终的结果.
	go func() {
		//for num := range ch {
		//	atomic.AddInt64(&sumVal, int64(sum(num)))
		//}
		for k := range nums {
			fmt.Println(k)
			atomic.AddInt64(&sumVal, int64(<-ch))
		}

		done <- struct{}{}
	}()

	for _, num := range nums {
		tmpNum := num
		go func() {
			ch <- sum(tmpNum)
		}()
	}
	// close(ch)

	select {
	case <-done:
		fmt.Println("handle complete")
		close(ch)
		return int(sumVal)
	}
}

func sum(a int) int {
	fmt.Printf("a: %d\n", a)
	return a + 10
}
