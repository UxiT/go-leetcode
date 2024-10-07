package algos

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p.Val != q.Val {
		return false
	}

	for p.Left != nil || q.Left != nil {
		p = p.Left
		q = q.Left

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}
	}

	return true
}
