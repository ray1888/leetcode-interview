package easy100

import "github.com/ray1888/leetcode-interview/datastructure"

func getMaxDepth(node *datastructure.TreeNode) int {
	if node == nil {
		return 0
	}
	return datastructure.Max(getMaxDepth(node.Left), getMaxDepth(node.Right)) + 1
}

func maxDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	return getMaxDepth(root)
}
