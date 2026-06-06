package linkedlists

/*
 English idea:
 
 Use two pointers, one starting at head1 and one starting at head2.
 Move both pointers one node at a time.
 When a pointer reaches the end of its list, move it to the start of the other list.
 This makes both pointers travel the same total distance.
 Eventually, they will meet at the same node in memory: the merge point.
 Return the data stored in that merge node.
*/

func findMergeNode(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) int32 {
	current1 := head1
	current2 := head2

	for current1 != current2 {
		if current1 == nil {
			current1 = head2
		} else {
			current1 = current1.next
		}

		if current2 == nil {
			current2 = head1
		} else {
			current2 = current2.next
		}
	}

	return current1.data
}
