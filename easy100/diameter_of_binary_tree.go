package easy100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func gmaxDepth(node *datastructure.TreeNode, result *[]int) int {
	if node == nil {
		return 0
	}
	left := gmaxDepth(node.Left, result)
	right := gmaxDepth(node.Right, result)
	if (*result)[0] < left+right {
		(*result)[0] = left + right
	}
	return datastructure.Max(left, right) + 1
}

func diameterOfBinaryTree(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	d := make([]int, 1)
	gmaxDepth(root, &d)
	return d[0]
}
