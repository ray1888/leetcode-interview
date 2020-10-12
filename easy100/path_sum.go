package easy100

import "github.com/ray1888/leetcode-interview/datastructure"

func dfsPathSum(node *datastructure.TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		if sum == node.Val {
			return true
		}
		return false
	}
	sum -= node.Val
	return dfsPathSum(node.Left, sum) || dfsPathSum(node.Right, sum)
}

func hasPathSum(root *datastructure.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return dfsPathSum(root, sum)
}
