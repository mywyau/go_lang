package linkedlists

/*
	insertNodeAtHead adds a new node to the beginning of a linked list.

	The current head of the list is passed in as `llist`.
	The value we want to insert is passed in as `data`.

	To insert at the head:
	1. Create a new node.
	2. Store `data` inside the new node.
	3. Point the new node's `next` field to the old head.
	4. Return the new node as the new head.

	Example:

	Current list:

		383 -> 484 -> nil

	Insert 321 at the head:

		321 -> 383 -> 484 -> nil

	This is O(1) time because we do not need to loop through the list.
*/

/*
	insertNodeAtHead adds a new node to the start of a linked list.
	
	Plain English idea:
	- Create a new node with the given data.
	- Point the new node's next to the current head.
	- Return the new node as the new head.
	
	Example:
	
		Current list: 20 -> 30 -> 40
		Insert 10 at head
		Result:       10 -> 20 -> 30 -> 40
*/

func insertNodeAtHead(llist *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{
		data: data,
		next: llist,
	}

	return newNode
}

// func insertNodeAtHead(llist * SinglyLinkedListNode, data int32) * SinglyLinkedListNode {
// 	newNode:= &SinglyLinkedListNode{
// 		data:data,
// 		next: llist,
// 	}
// }

func insertNodeAtHead_explicit(llist *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// Create the new node
	newNode := &SinglyLinkedListNode{}

	// Put the given value inside it
	newNode.data = data

	// Point the new node to the old head
	newNode.next = llist

	// Return the new node as the new head
	return newNode
}
