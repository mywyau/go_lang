package warmup

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	tests := []struct {
		name    string
		candles []int32
		want    int32
	}{
		{
			name:    "sample case",
			candles: []int32{3, 2, 1, 3},
			want:    2,
		},
		{
			name:    "all candles same height",
			candles: []int32{4, 4, 4, 4},
			want:    4,
		},
		{
			name:    "only one tallest candle",
			candles: []int32{1, 2, 3, 4},
			want:    1,
		},
		{
			name:    "tallest candle appears at start",
			candles: []int32{9, 1, 2, 9, 3},
			want:    2,
		},
		{
			name:    "single candle",
			candles: []int32{7},
			want:    1,
		},
		{
			name:    "larger repeated max",
			candles: []int32{1, 5, 5, 3, 5, 2},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := birthdayCakeCandles(tt.candles)

			if got != tt.want {
				t.Errorf("birthdayCakeCandles(%v) = %d; want %d", tt.candles, got, tt.want)
			}
		})
	}
}