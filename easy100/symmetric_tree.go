package easy100

func isSym(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil && node2 != nil {
		return false
	} else if node1 != nil && node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	if !isSym(node1.Left, node2.Right) {
		return false
	}
	if !isSym(node1.Right, node2.Left) {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSym(root.Left, root.Right)
}
