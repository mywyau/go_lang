package arrays

import (
	"reflect"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	tests := []struct {
		name    string
		n       int32
		queries [][]int32
		want    []int32
	}{
		{
			name: "HackerRank example",
			n:    2,
			queries: [][]int32{
				{1, 0, 5},
				{1, 1, 7},
				{1, 0, 3},
				{2, 1, 0},
				{2, 1, 1},
			},
			want: []int32{7, 3},
		},
		{
			name: "one sequence with multiple lookups",
			n:    1,
			queries: [][]int32{
				{1, 0, 5},
				{1, 0, 9},
				{2, 0, 0},
				{2, 0, 1},
			},
			want: []int32{5, 9},
		},
		{
			name: "lastAnswer changes the sequence index",
			n:    2,
			queries: [][]int32{
				{1, 0, 10}, // Add 10 to sequence 0
				{1, 1, 20}, // Add 20 to sequence 1
				{2, 0, 0},  // Read 10, so lastAnswer becomes 10
				{1, 1, 30}, // (1 XOR 10) % 2 = 1
				{2, 11, 1}, // (11 XOR 10) % 2 = 1; read 30
			},
			want: []int32{10, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dynamicArray(tt.n, tt.queries)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"dynamicArray(%d, %v) = %v, want %v",
					tt.n,
					tt.queries,
					got,
					tt.want,
				)
			}
		})
	}
}