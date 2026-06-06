package arrays

import "testing"

func TestHourglassSum(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int32
		want int32
	}{
		{
			name: "sample input returns 19",
			arr: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 19,
		},
		{
			name: "negative values example",
			arr: [][]int32{
				{-9, -9, -9, 1, 1, 1},
				{0, -9, 0, 4, 3, 2},
				{-9, -9, -9, 1, 2, 3},
				{0, 0, 8, 6, 6, 0},
				{0, 0, 0, -2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 28,
		},
		{
			name: "all negative numbers",
			arr: [][]int32{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
			},
			want: -7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hourglassSum(tt.arr)

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}