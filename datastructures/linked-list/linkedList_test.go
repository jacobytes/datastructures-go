package main

import "testing"

func TestLinkedList(t *testing.T) {

	t.Run("should instanciated Head as nil when creating LinkedList", func(t *testing.T) {
		// Act
		linkedList := LinkedList[int]{}

		// Assert
		if linkedList.Head != nil {
			t.Errorf("Head value wasn't nil")
		}
	})

	t.Run("should add Node as head when Append() is called on empty list", func(t *testing.T) {

		expected := 1
		// Arrange
		linkedList := LinkedList[int]{}

		// Act
		linkedList.Append(1)

		// Assert
		if linkedList.Head.Data != expected {
			t.Errorf("Head data value did not equal %v", expected)
		}

	})

	t.Run("should add nodes with next set as nil ", func(t *testing.T) {

		// Arrange
		linkedList := LinkedList[int]{}

		// Act
		linkedList.Append(1)

		// Assert
		if linkedList.Head.Next != nil {
			t.Errorf("Node was appended, but Next property wasn't nil.")
		}
	})

	t.Run("should set any extra nodes as Next property of any preceding Nodes", func(t *testing.T) {

		expected := 2
		// Arrange
		linkedList := LinkedList[int]{}
		linkedList.Append(1)

		// Act
		linkedList.Append(expected)

		// Assert
		if linkedList.Head.Next == nil {
			t.Errorf("No node appended, the Head's Next property is nil.")
		}
		if linkedList.Head.Next.Data != expected {
			t.Errorf("Node was appended, but Data property was %v instead of %v", linkedList.Head.Next.Data, expected)
		}
	})
}
