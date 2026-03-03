package linkedlist

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Reverse second half
	var prev *ListNode
	cur := slow.Next
	slow.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// 3. Merge two halves
	first, second := head, prev
	for second != nil {
		t1 := first.Next
		t2 := second.Next

		first.Next = second
		second.Next = t1

		first = t1
		second = t2
	}
}
