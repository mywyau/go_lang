package linkedlists

/*
Delete Node Notes

Example linked list:

20 -> 6 -> 2 -> 19 -> 7

Positions:

0     1    2    3     4

Deleting the head:
- If position == 0, return head.next
- This removes the first node

Deleting another node:
- Stop at the node before the one you want to delete
- Change its next pointer to skip over the deleted node

Example:

10 -> 20 -> 30 -> 40

Delete position 2, which is 30.

Stop at position 1, which is 20.

Then do:

current.next = current.next.next

Before:

20 -> 30 -> 40

After:

20 -> 40
*/

/*

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// If the list is empty, there is nothing to delete.
	if llist == nil {
		return nil
	}

	// If position is 0, we are deleting the head node.
	//
	// Example:
	//
	// 10 -> 20 -> 30
	//
	// Delete position 0:
	//
	// 20 -> 30
	//
	// So the new head becomes llist.next.
	if position == 0 {
		return llist.next
	}

	// We need to stop at the node BEFORE the node we want to delete.
	//
	// Example:
	//
	// 10 -> 20 -> 30 -> 40
	//
	// If we want to delete position 2, which is 30,
	// we need to stop at position 1, which is 20.
	current := llist

	for i := int32(0); i < position-1; i++ {
		current = current.next
	}

	// current is now the node before the one we want to delete.
	//
	// current.next is the node we want to remove.
	//
	// So we skip over it:
	//
	// Before:
	//
	// current -> nodeToDelete -> nextNode
	//
	// After:
	//
	// current -> nextNode
	current.next = current.next.next

	return llist
}

*/


// deleteNode removes a node from a linked list at the given position.
//
// Plain English idea:
// - If position is 0, delete the head by returning llist.next.
// - Otherwise, walk to the node before the one we want to delete.
// - Change current.next to skip over the deleted node.
// - Return the original head.
//
// Example:
//
//	10 -> 20 -> 30 -> 40
//
// Delete position 2:
//
//	10 -> 20 -> 40

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if position == 0 {
		return llist.next
	}

	current := llist

	for i := int32(0); i < position-1; i++ {
		current = current.next
	}

	current.next = current.next.next

	return llist
}


