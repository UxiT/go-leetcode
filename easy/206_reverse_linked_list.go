package easy

func ReverseList(head *ListNode) *ListNode {
	var (
		prev, next *ListNode
	)

	for head != nil {
		next = head.Next

		head.Next = prev
		prev = head

		head = next
	}

	return prev
}
