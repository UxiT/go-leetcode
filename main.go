package main

import (
	"fmt"
	"leetcode/easy"
	"leetcode/medium"
)

func main() {
	nums := []int{5, 1, 2, 3, 4}
	res := medium.FindMin(nums)

	fmt.Println(res)
}

func buildList(nums []int) *easy.ListNode {
	list := &easy.ListNode{}

	if len(nums) == 1 {
		return &easy.ListNode{Val: nums[0]}
	}

	for i := 0; i < len(nums); i++ {
		list.Val = nums[i]
		list.Next = buildList(nums[:i])
	}

	return list
}
