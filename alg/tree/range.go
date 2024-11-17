package tree

import "container/list"

// Node 二叉树的节点.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// zigzagLevelOrder 二叉树的锯齿形遍历.
func zigzagLevelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	stack := list.New()
	stack.PushBack(root) // 先将根节点入栈.

	// 对栈进行出栈操作.
	var i = 0
	var e *Node
	for stack.Len() != 0 { // 栈内有数据，则该层不为空.
		// 将该层的数据全部弹出，边弹出数据，边将左右节点入栈.
		tmpResult := make([]int, 0, stack.Len())
		ie := stack.Back()
		e, _ = ie.Value.(*Node)
		if e != nil {
			tmpResult = append(tmpResult, e.Val)
		}

		// 会增加用一个临时栈，将下一层节点数据全部入栈.

		if i/2 == 0 { // 偶数层，从左向右遍历.

		}

		i++
	}

}
