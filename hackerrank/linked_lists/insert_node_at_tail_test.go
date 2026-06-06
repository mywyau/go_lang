package linkedlists

import (
	"reflect"
	"testing"
)

func listToSliceForTailTest(head *SinglyLinkedListNode) []int32 {
	var result []int32

	current := head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}

	return result
}

func buildListForTailTest(values []int32) *SinglyLinkedListNode {
	var head *SinglyLinkedListNode

	for _, value := range values {
		head = insertNodeAtTail(head, value)
	}

	return head
}

func TestInsertNodeAtTail(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		data  int32
		want  []int32
	}{
		{
			name:  "insert into empty list",
			input: []int32{},
			data:  10,
			want:  []int32{10},
		},
		{
			name:  "insert into one node list",
			input: []int32{10},
			data:  20,
			want:  []int32{10, 20},
		},
		{
			name:  "insert into multiple node list",
			input: []int32{10, 20, 30},
			data:  40,
			want:  []int32{10, 20, 30, 40},
		},
		{
			name:  "insert negative number",
			input: []int32{1, 2, 3},
			data:  -1,
			want:  []int32{1, 2, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListForTailTest(tt.input)

			gotHead := insertNodeAtTail(head, tt.data)

			got := listToSliceForTailTest(gotHead)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}