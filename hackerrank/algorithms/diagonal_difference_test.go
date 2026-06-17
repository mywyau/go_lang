package algorithms

import "testing"

func TestDiagonalDifference(t *testing.T) {

	tests := []struct {
		name string
		arr  [][]int32
		want int32
	}{
		{
			name: "sample input",
			arr: [][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			want: 15,
		},
		{
			name: "diagonals are equal",
			arr: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 0,
		},
		{
			name: "single element matrix",
			arr: [][]int32{
				{5},
			},
			want: 0,
		},
		{
			name: "two by two matrix",
			arr: [][]int32{
				{1, 2},
				{3, 4},
			},
			want: 0,
		},
		{
			name: "negative numbers",
			arr: [][]int32{
				{-1, -2, -3},
				{-4, -5, -6},
				{-7, -8, -9},
			},
			want: 0,
		},
		{
			name: "larger difference",
			arr: [][]int32{
				{10, 2, 3},
				{4, 20, 6},
				{7, 8, 30},
			},
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diagonalDifference(tt.arr)

			if got != tt.want {
				t.Errorf("diagonalDifference() = %d, want %d", got, tt.want)
			}
		})
	}
}