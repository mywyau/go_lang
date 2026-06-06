package linkedlists

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		position int32
		want     []int32
	}{
		{
			name:     "delete head node",
			input:    []int32{20, 6, 2, 19, 7},
			position: 0,
			want:     []int32{6, 2, 19, 7},
		},
		{
			name:     "delete middle node",
			input:    []int32{20, 6, 2, 19, 7},
			position: 2,
			want:     []int32{20, 6, 19, 7},
		},
		{
			name:     "delete last node",
			input:    []int32{20, 6, 2, 19, 7},
			position: 4,
			want:     []int32{20, 6, 2, 19},
		},
		{
			name:     "delete from two node list",
			input:    []int32{10, 20},
			position: 1,
			want:     []int32{10},
		},
		{
			name:     "delete only node",
			input:    []int32{10},
			position: 0,
			want:     []int32(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildLinkedList(tt.input)

			gotHead := deleteNode(head, tt.position)

			got := linkedListToSlice(gotHead)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
