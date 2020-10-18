package medium100

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *datastructure.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	stack := make([]*datastructure.TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		length := len(stack)
		elem := make([]int, 0)
		for i := 0; i < length; i++ {
			node := stack[i]
			elem = append(elem, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, elem)
		stack = stack[length:]
	}
	return result
}
