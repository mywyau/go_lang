package linkedlists

// insertNodeAtTail adds a new node to the end of a linked list.
//
// Plain English idea:
// - Create a new node with the given data.
// - If the list is empty, the new node becomes the head.
// - Otherwise, walk through the list until we reach the last node.
// - Point the last node's next to the new node.
// - Return the original head.
//
// Example:
//
//	Current list: 10 -> 20 -> 30
//	Insert 40 at tail
//	Result:       10 -> 20 -> 30 -> 40

func insertNodeAtTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{data: data}

	if head == nil {
		return newNode
	}

	current := head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	return head
}