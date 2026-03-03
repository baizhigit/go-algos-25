package linkedlist

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. reverse second half
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 3. compare halves
	left := head
	right := prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}

	return true
}
