package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]

		lists = lists[2:]

		merged := merge(l1, l2)
		lists = append(lists, merged)
	}

	return lists[0]
}

func merge(list1, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	current := merged

	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	return merged.Next
}
