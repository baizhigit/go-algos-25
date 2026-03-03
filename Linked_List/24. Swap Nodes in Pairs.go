package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		left := current.Next
		right := current.Next.Next

		left.Next = right.Next
		right.Next = left
		current.Next = right
		current = left
	}

	return dummy.Next
}
