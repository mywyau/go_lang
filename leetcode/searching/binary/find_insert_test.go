package binary

import "testing"

func TestFindInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "target exists in middle",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "target should be inserted in middle",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "target should be inserted at end",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "target should be inserted at start",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "single element target exists",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "single element insert before",
			nums:   []int{1},
			target: 0,
			want:   0,
		},
		{
			name:   "single element insert after",
			nums:   []int{1},
			target: 2,
			want:   1,
		},
		{
			name:   "empty slice",
			nums:   []int{},
			target: 5,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindInsert(tt.nums, tt.target)

			if got != tt.want {
				t.Errorf("FindInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
