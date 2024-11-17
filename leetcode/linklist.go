package leetcode

import "errors"

/*
给你一个链表的头节点 head 。删除 链表的 中间节点 ，并返回修改后的链表的头节点 head 。
长度为 n 链表的中间节点是从头数起第 ⌊n / 2⌋ 个节点（下标从 0 开始），其中 ⌊x⌋ 表示小于或等于 x 的最大整数。

示例：
输入：head = [1,3,4,7,1,2,6]
输出：[1,3,4,1,2,6]
*/

// ListNode 链表节点.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) len() int {
	var count int
	for {
		if l.Next == nil {
			break
		}

		count++
	}
	return count
}

// deleteMiddle 删除链表的中间节点.
func deleteMiddle(head *ListNode) *ListNode {
	delIndex := head.len() / 2
	var currentNode *ListNode
	var preNode *ListNode
	for i := 0; i < delIndex; i++ {
		if i == 0 {
			currentNode = head
			preNode = head
		} else {
			preNode = currentNode
			currentNode = currentNode.Next
		}
	}

	// 将 currentNode 的 next 阶段指向其前一个节点.
	preNode.Next = currentNode.Next
	return head
}

// reverseList 反转链表.
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	currentNode := head.Next
	nextNode := head.Next.Next
	preNode := head
	for {
		// 将指针往前指.
		currentNode.Next = preNode // 将当前节点的 next 指向前一个.
		preNode = currentNode
		currentNode = currentNode.Next
		nextNode = currentNode.Next // 下一个节点的下一个节点.
		if nextNode.Next == nil {
			break
		}
	}

	return nextNode
}

/*
实现一个 lru 缓存.
*/
type cacheNode struct {
	data int
	key  string
	next *cacheNode
}

type lruCache struct {
	cap       int        // 容量.
	nodeCount int        // 已存入数据量.
	head      *cacheNode // 第一个节点，不存储数据，仅用于指向缓存链表.
}

func NewLruCache(cap int) *lruCache {
	return &lruCache{
		cap:  cap,
		head: &cacheNode{},
	}
}

func (c *lruCache) moveTailFunc(node *cacheNode) {
	// 找到最后一个节点.
	var lastNode *cacheNode
	currentNode := c.head
	for {
		if currentNode.next == nil {
			lastNode = currentNode
			break
		}
	}

	lastNode.next = node
	node.next = nil
}

func (c *lruCache) Get(key string) (int, error) {
	if c.nodeCount == 0 {
		return 0, errors.New("cache is empty")
	}

	node := c.head
	var preNode *cacheNode
	for {
		if node == nil { // 已经找过最后一个节点了.
			break
		}

		if node.key == key { // 找到这个节点后，将该节点移动到链表的最后.
			tmpNode := node          // 将当前节点赋值给一个临时节点，删除该节点.
			preNode.next = node.next // 将前一个节点的 next 指向当前节点的 next.
			c.moveTailFunc(tmpNode)
			return node.data, nil
		}
		preNode = node // 赋值前一个节点.
		node = node.next
	}
	return 0, errors.New("cache not found this key")
}

func (c *lruCache) Set(key string, data int) error {
	// 判断是否已满.
	if c.nodeCount == c.cap {
		// 当缓存已满时，移除节点头的数据.
		tmpNode := c.head
		c.head.next = tmpNode.next.next // 直接将头结点的 next 指针，指向原 head 节点指向的第一个节点的下一个节点.
		c.nodeCount--
	}

	// 添加一个数据进去.
	newNode := &cacheNode{
		data: data,
		key:  key,
		next: nil,
	}
	c.moveTailFunc(newNode)
	c.nodeCount++
	return nil
}
