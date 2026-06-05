package linkedlists

func compareLists(llist1 *SinglyLinkedListNode, llist2 *SinglyLinkedListNode) bool {
	for llist1 != nil && llist2 != nil {
		if llist1.data != llist2.data {
			return false
		}

		llist1 = llist1.next
		llist2 = llist2.next
	}

	return llist1 == nil && llist2 == nil
}


