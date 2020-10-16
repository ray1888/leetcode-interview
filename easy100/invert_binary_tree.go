package easy100

import "github.com/ray1888/leetcode-interview/datastructure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invert(node *datastructure.TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		tmp := node.Right
		node.Right = node.Left
		node.Left = tmp
	} else if node.Left != nil && node.Right == nil {
		node.Right = node.Left
		node.Left = nil
	} else if node.Left == nil && node.Right != nil {
		node.Left = node.Right
		node.Right = nil
	}
	invert(node.Left)
	invert(node.Right)
}

func invertTree(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	invert(root)
	return root
}
