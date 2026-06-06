package linkedlists

import "testing"

func TestCompareLinkedLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int32
		list2 []int32
		want  bool
	}{
		{
			name:  "same linked lists",
			list1: []int32{20, 6, 2, 19, 7},
			list2: []int32{20, 6, 2, 19, 7},
			want:  true,
		},
		{
			name:  "different value",
			list1: []int32{20, 6, 2, 19, 7},
			list2: []int32{20, 6, 2, 99, 7},
			want:  false,
		},
		{
			name:  "first list longer",
			list1: []int32{20, 6, 2, 19, 7},
			list2: []int32{20, 6, 2},
			want:  false,
		},
		{
			name:  "second list longer",
			list1: []int32{20, 6},
			list2: []int32{20, 6, 2},
			want:  false,
		},
		{
			name:  "both lists empty",
			list1: []int32{},
			list2: []int32{},
			want:  true,
		},
		{
			name:  "one list empty",
			list1: []int32{10},
			list2: []int32{},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head1 := buildLinkedList(tt.list1)
			head2 := buildLinkedList(tt.list2)

			got := compareLinkedLists(head1, head2)

			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}