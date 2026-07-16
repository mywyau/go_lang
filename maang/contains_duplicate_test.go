package maang

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name  string
		given []int
		want  bool
	}{
		{
			name:  "contains duplicate",
			given: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			name:  "all values are unique",
			given: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			name:  "contains several duplicates",
			given: []int{1, 1, 1, 3, 3, 4},
			want:  true,
		},
		{
			name:  "empty slice",
			given: []int{},
			want:  false,
		},
		{
			name:  "single value",
			given: []int{10},
			want:  false,
		},
		{
			name:  "duplicate negative value",
			given: []int{-1, 2, -1},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.given)

			if got != tt.want {
				t.Errorf(
					"containsDuplicate(%v) = %v, want %v",
					tt.given,
					got,
					tt.want,
				)
			}
		})
	}
}
