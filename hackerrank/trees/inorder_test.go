package trees

import (
	"testing"
)

func TestInOrder(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   string
	}{
		{
			name:   "example tree",
			values: []int{1, 2, 5, 3, 6, 4},
			want:   "1 2 3 4 5 6 ",
		},
		{
			name:   "single node",
			values: []int{10},
			want:   "10 ",
		},
		{
			name:   "empty tree",
			values: []int{},
			want:   "",
		},
		{
			name:   "left side tree",
			values: []int{5, 4, 3, 2, 1},
			want:   "1 2 3 4 5 ",
		},
		{
			name:   "right side tree",
			values: []int{1, 2, 3, 4, 5},
			want:   "1 2 3 4 5 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode

			for _, value := range tt.values {
				root = InsertTreeNode(root, value)
			}

			got := captureOutput(func() {
				inorder(root)
			})

			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
