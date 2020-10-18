package medium100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func dfsTreeInorder(node *datastructure.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfsTreeInorder(node.Left, result)
	*result = append(*result, node.Val)
	dfsTreeInorder(node.Right, result)
}

func inorderTraversal(root *datastructure.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	dfsTreeInorder(root, &result)
	return result
}
