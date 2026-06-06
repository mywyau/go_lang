package linkedlists

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// func printLinkedList(head *SinglyLinkedListNode) {
// 	// Start at the first node in the linked list
// 	current := head

// 	// Keep looping while current is not nil
// 	// nil means we have reached the end of the linked list
// 	for current != nil {
// 		// Print the value stored in the current node
// 		fmt.Println(current.data)

// 		// Move to the next node
// 		current = current.next
// 	}
// }

func printLinkedList(head *SinglyLinkedListNode) {
	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func printLinkedList2(head *SinglyLinkedListNode) {
	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}