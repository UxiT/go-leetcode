package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var result []int
	var goDeeper func(node *TreeNode)

	goDeeper = func(node *TreeNode) {
		if node == nil {
			return
		}

		goDeeper(node.Left)
		result = append(result, node.Val)
		goDeeper(node.Right)
	}

	goDeeper(root)
	return result
}
