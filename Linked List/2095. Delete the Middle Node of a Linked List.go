package linkedlist

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
