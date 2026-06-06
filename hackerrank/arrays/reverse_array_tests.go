package arrays

import "testing"

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name string
		input []int32
		want []int32
	}{
		{
			name:  "reverses normal array",
			input: []int32{1, 4, 3, 2},
			want:  []int32{2, 3, 4, 1},
		},
		{
			name:  "reverses already sorted array",
			input: []int32{1, 2, 3, 4, 5},
			want:  []int32{5, 4, 3, 2, 1},
		},
		{
			name:  "single item array",
			input: []int32{10},
			want:  []int32{10},
		},
		{
			name:  "empty array",
			input: []int32{},
			want:  []int32{},
		},
		{
			name:  "array with negative numbers",
			input: []int32{-1, -2, -3},
			want:  []int32{-3, -2, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseArray(tt.input)

			if !equalSlices(got, tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func equalSlices(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}