package leetcode

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
方法一：深度优先搜索
如果我们知道了左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为
max(l,r)+1
而左子树和右子树的最大深度又可以以同样的方式进行计算。因此我们可以用「深度优先搜索」的方法来计算二叉树的最大深度。
具体而言，在计算当前二叉树的最大深度时，可以先递归计算出其左子树和右子树的最大深度，然后在 O(1) 时间内计算出当前二叉树的最大深度。
递归在访问到空节点时退出。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := maxDepth(root.Left)   // 求左子树深度.
	rightDepth := maxDepth(root.Right) // 求右子树深度.
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

/*
给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。

输入：root = [3,1,4,3,null,1,5]
输出：4
解释：图中蓝色节点为好节点。
根节点 (3) 永远是个好节点。
节点 4 -> (3,4) 是路径中的最大值。
节点 5 -> (3,4,5) 是路径中的最大值。
节点 3 -> (3,1,3) 是路径中的最大值。
*/
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 左右子树分别找好节点.
	var leftCount, rightCount int
	leftNode := root.Left
	leftMaxVal := root.Val
	for leftNode.Left != nil {
		if leftNode.Val >= leftMaxVal {
			leftCount++
			leftMaxVal = leftNode.Val
		}
		leftNode = leftNode.Left
	}

	rightNode := root.Right
	rightMaxVal := root.Val
	for rightNode.Right != nil {
		if rightNode.Val >= rightMaxVal {
			rightCount++
			leftMaxVal = rightNode.Val
		}
		rightNode = rightNode.Right
	}

	return leftCount + rightCount + 1
}
