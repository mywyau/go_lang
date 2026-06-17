package algorithms

import "testing"

func TestCalculateRatios(t *testing.T) {
	tests := []struct {
		name         string
		arr          []int32
		wantPositive float64
		wantNegative float64
		wantZero     float64
	}{
		{
			name:         "sample input",
			arr:          []int32{-4, 3, -9, 0, 4, 1},
			wantPositive: 0.5,
			wantNegative: 0.3333333333333333,
			wantZero:     0.16666666666666666,
		},
		{
			name:         "all positive",
			arr:          []int32{1, 2, 3, 4},
			wantPositive: 1.0,
			wantNegative: 0.0,
			wantZero:     0.0,
		},
		{
			name:         "all negative",
			arr:          []int32{-1, -2, -3, -4},
			wantPositive: 0.0,
			wantNegative: 1.0,
			wantZero:     0.0,
		},
		{
			name:         "all zeroes",
			arr:          []int32{0, 0, 0, 0},
			wantPositive: 0.0,
			wantNegative: 0.0,
			wantZero:     1.0,
		},
		{
			name:         "mixed values",
			arr:          []int32{-1, 0, 1},
			wantPositive: 0.3333333333333333,
			wantNegative: 0.3333333333333333,
			wantZero:     0.3333333333333333,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPositive, gotNegative, gotZero := calculateRatios(tt.arr)

			if gotPositive != tt.wantPositive {
				t.Errorf("positive ratio = %f, want %f", gotPositive, tt.wantPositive)
			}

			if gotNegative != tt.wantNegative {
				t.Errorf("negative ratio = %f, want %f", gotNegative, tt.wantNegative)
			}

			if gotZero != tt.wantZero {
				t.Errorf("zero ratio = %f, want %f", gotZero, tt.wantZero)
			}
		})
	}
}