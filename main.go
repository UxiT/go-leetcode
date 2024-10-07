package main

import (
	"fmt"
	algos "leetcode/easy"
)

func main() {

	node2 := algos.TreeNode{Val: 2, Left: nil, Right: nil}
	root1 := algos.TreeNode{Val: 1, Left: &node2, Right: nil}
	root2 := algos.TreeNode{Val: 1, Left: nil, Right: &node2}

	res := algos.IsSameTree(&root1, &root2)

	fmt.Printf("res: %v", res)
}
