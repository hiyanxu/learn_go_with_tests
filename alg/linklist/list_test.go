package linklist

import "testing"

func TestLinkList(t *testing.T) {
	l := initLinkList(1)
	l.AppendNode(2)
	l.AppendNode(3)
	l.AppendNode(4)
	l.AppendNode(5)

	// 遍历输出链表中的所有节点.
	l.Range(l.HeadNode)
}
