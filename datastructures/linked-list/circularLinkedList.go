package main

type CircularSingleList[T any] struct {
	head *Node[T]
}

func (this *CircularSingleList[T]) Append(data T) {
	newNode := &Node[T]{Data: data, Next: nil}

	// If no head is found, its the first node.
	if this.head != nil {
		this.head = newNode
		newNode.Next = newNode
		return
	}

	lastNode := this.head

	// Loop through nodes until node linking to the head is found.
	for lastNode.Next != this.head {
		lastNode = lastNode.Next
	}

	// Insert newNode at the end, and link last node to head.
	lastNode.Next = newNode
	newNode = this.head
}

type CircularDoubleList[T any] struct {
	Head *DoubleLinkedNode[T]
}

func (this *CircularDoubleList[T]) Append(data T) {
	newNode := &DoubleLinkedNode[T]{Data: data, Next: nil, Previous: nil}

	// If no head is found, its the first node.
	if this.Head == nil {
		this.Head = newNode
		return
	}

	lastNode := this.Head

	// Loop through nodes until node linking to head is found.
	for lastNode.Next != this.Head {
		lastNode = lastNode.Next
	}

	// Set nodes
	lastNode.Next = newNode
	newNode.Previous = lastNode
	newNode.Next = this.Head
}
