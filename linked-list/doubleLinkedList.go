package main

type DoubleLinkedNode[T any] struct {
	Data     T
	Next     *DoubleLinkedNode[T]
	Previous *DoubleLinkedNode[T]
}

type DoubleLinkedList[T any] struct {
	Head *DoubleLinkedNode[T]
}

func (this *DoubleLinkedList[T]) Append(data T) {
	newNode := &DoubleLinkedNode[T]{Data: data, Next: nil, Previous: nil}

	// If no head is found, its the first node.
	if this.Head == nil {
		this.Head = newNode
		return
	}

	lastNode := this.Head

	// Loop through nodes until node with no next node is found.
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	// Set nodes
	lastNode.Next = newNode
	newNode.Previous = lastNode
}
