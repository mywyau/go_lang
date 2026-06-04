package arrays

import "testing"

func TestSimpleArraySum(t *testing.T) {
	tests := []struct {
		name string
		ar   []int32
		want int32
	}{
		{
			name: "sample input",
			ar:   []int32{1, 2, 3, 4, 10, 11},
			want: 31,
		},
		{
			name: "empty array",
			ar:   []int32{},
			want: 0,
		},
		{
			name: "single value",
			ar:   []int32{5},
			want: 5,
		},
		{
			name: "negative numbers",
			ar:   []int32{-1, -2, -3},
			want: -6,
		},
		{
			name: "mixed positive and negative numbers",
			ar:   []int32{-1, 2, -3, 4},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simpleArraySum(tt.ar)

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}