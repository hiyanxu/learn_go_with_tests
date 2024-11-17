package linklist

import "fmt"

// Node 链表节点.
type Node struct {
	data int
	next *Node
}

// LinkList 链表.
type LinkList struct {
	HeadNode *Node // 具有一个头节点.
}

func initLinkList(data int) *LinkList {
	return &LinkList{
		HeadNode: &Node{
			data: data,
			next: nil,
		}}
}

// findNilNode 找到最后一个 nil 的节点，表示待插入.
func (l *LinkList) findNilNode(n *Node) *Node {
	if n.next == nil {
		return n
	} else {
		return l.findNilNode(n.next)
	}
}

// AppendNode 追加一个节点.
func (l *LinkList) AppendNode(data int) {
	node := &Node{
		data: data,
		next: nil,
	}

	// 找到待插入的节点.
	lastNode := l.findNilNode(l.HeadNode)
	lastNode.next = node
}

func (l *LinkList) Range(n *Node) {
	if n.next == nil {
		fmt.Println(n.data)
		fmt.Println("结束了")
		return
	} else {
		fmt.Println(n.data)
		l.Range(n.next)
	}
}

// ReverseList 反转链接.
func (l *LinkList) ReverseList() {
	preNode := l.HeadNode // 当前节点.
	tmpNode := l.HeadNode // 临时节点，保存下一位节点的地址.
	l.HeadNode.next = nil // 将首位节点置为最后一个节点.
	for tmpNode.next != nil {
		tmpNode.next = tmpNode
		tmpNode.next.next = preNode // 将下一个节点反指回当前节点.
		preNode = tmpNode

	}
}

// mergeTwoList 合并两个有序链表.
func (l *LinkList) mergeTwoList(l2 *LinkList) *LinkList {
	var resp *LinkList
	var lCurNode *Node
	var l2CurNode *Node
	if l.HeadNode.data > l2.HeadNode.data {
		l2CurNode = l2.HeadNode.next
		resp = initLinkList(l2.HeadNode.data)
	} else {
		lCurNode = l.HeadNode.next
		resp = initLinkList(l.HeadNode.data)
	}

	for lCurNode.next == nil || l2CurNode.next == nil {

	}
}
