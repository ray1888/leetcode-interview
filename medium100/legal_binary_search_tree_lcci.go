package medium100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func getMin(node *datastructure.TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

func getMax(node *datastructure.TreeNode) int {
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}

func isValidBST(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	leftValid := root.Left == nil || (getMax(root.Left) < root.Val)
	rightValid := root.Right == nil || (getMin(root.Right) > root.Val)
	return leftValid && rightValid && isValidBST(root.Left) && isValidBST(root.Right)
}

func isValidBSTBound(node, lower, top *datastructure.TreeNode) bool {
	if node == nil {
		return true
	}
	if lower != nil && lower.Val >= node.Val {
		return false
	}
	if top != nil && top.Val <= node.Val {
		return false
	}
	return isValidBSTBound(node.Left, lower, node) && isValidBSTBound(node.Right, node, top)
}

func isValidBSTBoundOuter(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTBound(root, nil, nil)
}
