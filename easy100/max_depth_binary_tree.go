package easy100

import "github.com/ray1888/leetcode-interview/datastructure"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMaxDepth(node *datastructure.TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getMaxDepth(node.Left), getMaxDepth(node.Right)) + 1
}

func maxDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	return getMaxDepth(root)
}
