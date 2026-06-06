package linkedlists

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