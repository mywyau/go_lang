package algorithms

import "testing"

func TestCalculateMiniMaxSum(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int32
		wantMin int64
		wantMax int64
	}{
		{
			name:    "sample input",
			arr:     []int32{1, 2, 3, 4, 5},
			wantMin: 10,
			wantMax: 14,
		},
		{
			name:    "hackerrank sample with unsorted values",
			arr:     []int32{7, 69, 2, 221, 8974},
			wantMin: 299,
			wantMax: 9271,
		},
		{
			name:    "all same values",
			arr:     []int32{5, 5, 5, 5, 5},
			wantMin: 20,
			wantMax: 20,
		},
		{
			name:    "large int32 values",
			arr:     []int32{2147483647, 2147483647, 2147483647, 2147483647, 2147483647},
			wantMin: 8589934588,
			wantMax: 8589934588,
		},
		{
			name:    "mixed small values",
			arr:     []int32{5, 1, 9, 3, 7},
			wantMin: 16,
			wantMax: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := calculateMiniMaxSum(tt.arr)

			if gotMin != tt.wantMin {
				t.Errorf("min sum = %d, want %d", gotMin, tt.wantMin)
			}

			if gotMax != tt.wantMax {
				t.Errorf("max sum = %d, want %d", gotMax, tt.wantMax)
			}
		})
	}
}