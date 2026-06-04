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

func deleteNode2(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {

	if position == 0 {
		return llist.next
	}

	// current := llist  // short form type inference
	var current *SinglyLinkedListNode = llist // long form

	// i < position - 1 is the node we want to stop before the node we wish to delete
	// we go from 0 -> position - 1, i++ to move all the pointers along for current
	for i := int32(0); i < position-1; i++ {
		current = current.next
	}

	current.next = current.next.next // point the current.next to the node after we want to delete, effectively deleting it

	return llist
}
