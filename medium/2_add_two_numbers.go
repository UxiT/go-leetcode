package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy

	for l1 != nil || l2 != nil {
		if l1 != nil {
			node.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			node.Val += l2.Val
			l2 = l2.Next
		}

		if node.Val > 9 {
			node.Val -= 10
			node.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			node.Next = &ListNode{}
		}

		node = node.Next
	}

	return dummy
}
