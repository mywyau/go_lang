package trees

import "testing"

func TestGetHeight(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "example tree",
			values: []int{1, 2, 5, 3, 6, 4},
			want:   3,
		},
		{
			name:   "single node",
			values: []int{10},
			want:   0,
		},
		{
			name:   "empty tree",
			values: []int{},
			want:   -1,
		},
		{
			name:   "left side tree",
			values: []int{5, 4, 3, 2, 1},
			want:   4,
		},
		{
			name:   "right side tree",
			values: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "balanced tree",
			values: []int{4, 2, 6, 1, 3, 5, 7},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode

			for _, value := range tt.values {
				root = InsertTreeNode(root, value)
			}

			got := getHeight(root)

			if got != tt.want {
				t.Fatalf("expected height %d, got %d", tt.want, got)
			}
		})
	}
}