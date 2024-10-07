package easy

func IsSymmetric(root *TreeNode) bool {
	var traverse func(nodeL, nodeR *TreeNode)
	result := true

	traverse = func(nodeL, nodeR *TreeNode) {
		if (nodeL == nil && nodeR != nil) || (nodeL != nil && nodeR == nil) {
			result = false
			return
		}

		if nodeL == nil && nodeR == nil {
			return
		}

		traverse(nodeL.Left, nodeR.Right)
		if nodeL.Val != nodeR.Val {
			result = false
			return
		}
		traverse(nodeL.Right, nodeR.Left)
	}

	traverse(root, root)

	return result
}
