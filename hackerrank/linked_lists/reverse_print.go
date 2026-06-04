package linkedlists

import "fmt"

func reversePrint(llist *SinglyLinkedListNode) {
	if llist == nil {
		return
	}

	reversePrint(llist.next)

	fmt.Println(llist.data)
}

