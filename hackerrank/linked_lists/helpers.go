package linkedlists

func buildLinkedList(values []int32) *SinglyLinkedListNode {
	if len(values) == 0 {
		return nil
	}

	head := &SinglyLinkedListNode{data: values[0]}
	current := head

	for _, value := range values[1:] {
		current.next = &SinglyLinkedListNode{data: value}
		current = current.next
	}

	return head
}

func linkedListToSlice(head *SinglyLinkedListNode) []int32 {
	var result []int32

	current := head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}

	return result
}
