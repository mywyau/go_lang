package dynamicprogramming

import "testing"

func TestGetWays(t *testing.T) {
	tests := []struct {
		name   string
		amount int32
		coins  []int32
		want   int64
	}{
		{
			name:   "example amount 4 with coins 1 2 3",
			amount: 4,
			coins:  []int32{1, 2, 3},
			want:   4,
		},
		{
			name:   "example amount 10 with coins 2 5 3 6",
			amount: 10,
			coins:  []int32{2, 5, 3, 6},
			want:   5,
		},
		{
			name:   "amount 0 has one way using no coins",
			amount: 0,
			coins:  []int32{1, 2, 3},
			want:   1,
		},
		{
			name:   "cannot make odd amount with only coin 2",
			amount: 3,
			coins:  []int32{2},
			want:   0,
		},
		{
			name:   "single coin exactly matches amount",
			amount: 5,
			coins:  []int32{5},
			want:   1,
		},
		{
			name:   "single coin can build amount repeatedly",
			amount: 6,
			coins:  []int32{2},
			want:   1,
		},
		{
			name:   "coin values larger than amount",
			amount: 3,
			coins:  []int32{5, 10},
			want:   0,
		},
		{
			name:   "amount 5 with coins 1 2 5",
			amount: 5,
			coins:  []int32{1, 2, 5},
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWays(tt.amount, tt.coins)

			if got != tt.want {
				t.Errorf("getWays(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
			}
		})
	}
}