package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy

	// move left n+1 steps ahead
	for i := 0; i <= n; i++ {
		left = left.Next
	}

	// move both until left hits end
	for left != nil {
		left = left.Next
		right = right.Next
	}

	right.Next = right.Next.Next
	return dummy.Next
}
