package easy100

import (
	"math"

	"github.com/ray1888/leetcode-interview/datastructure"
)

func getMinimumDifference(root *datastructure.TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*datastructure.TreeNode)
	dfs = func(node *datastructure.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
