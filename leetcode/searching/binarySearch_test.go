package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "finds target in middle",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "finds target at start",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "finds target at end",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			name:   "returns -1 when target is not found",
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   -1,
		},
		{
			name:   "returns -1 for empty slice",
			nums:   []int{},
			target: 1,
			want:   -1,
		},
		{
			name:   "works with one item found",
			nums:   []int{7},
			target: 7,
			want:   0,
		},
		{
			name:   "works with one item not found",
			nums:   []int{7},
			target: 3,
			want:   -1,
		},
		{
			name:   "works with negative numbers",
			nums:   []int{-10, -5, 0, 5, 10},
			target: -5,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.nums, tt.target)

			if got != tt.want {
				t.Fatalf("BinarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
