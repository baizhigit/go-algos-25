package linkedlist

// Doubly Linked List with Head, Tail & Size

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	if this.Head == nil {
		this.Head = node
		this.Tail = node
	} else {
		node.Next = this.Head
		this.Head.Prev = node
		this.Head = node
	}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if this.Tail == nil {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.Next = node
		node.Prev = this.Tail
		this.Tail = node
	}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	node := &Node{Val: val}
	prev := curr.Prev

	prev.Next = node
	node.Prev = prev
	node.Next = curr
	curr.Prev = node

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Head != nil {
			this.Head.Prev = nil
		} else {
			this.Tail = nil
		}
	} else if index == this.Size-1 {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
	} else {
		curr := this.Head
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}
	this.Size--
}
