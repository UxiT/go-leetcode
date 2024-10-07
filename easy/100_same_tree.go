package algos

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	var goDeeper func(node1 *TreeNode, node2 *TreeNode)
	result := true

	goDeeper = func(node1 *TreeNode, node2 *TreeNode) {
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			result = false
			return
		}

		if node1 == nil && node2 == nil {
			return
		}

		goDeeper(node1.Left, node2.Left)
		if node1.Val != node2.Val {
			result = false
			return
		}
		goDeeper(node1.Right, node2.Right)
	}

	goDeeper(p, q)

	return result
}
