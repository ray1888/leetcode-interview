package medium100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"github.com/ray1888/leetcode-interview/datastructure"
)

func reverse(array *[]int) {
	i := 0
	j := len(*array) - 1
	for i < j {
		tmp := (*array)[i]
		(*array)[i] = (*array)[j]
		(*array)[j] = tmp
		i++
		j--
	}

}

func zigzagLevelOrder(root *datastructure.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	stack := make([]*datastructure.TreeNode, 0)
	stack = append(stack, root)
	count := 0
	for len(stack) > 0 {
		count++
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
		if count%2 == 0 {
			reverse(&elem)
		}
		result = append(result, elem)
		stack = stack[length:]
	}
	return result
}
