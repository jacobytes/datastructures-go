package main

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func (this *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{Data: data, Next: nil}

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

	// Set newNode as next node on the last node
	lastNode.Next = newNode
}

func (this *LinkedList[T]) Remove(data T) {

}
