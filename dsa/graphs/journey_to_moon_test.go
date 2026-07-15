package graphs

import "testing"

func TestJourneyToMoon(t *testing.T) {
	tests := []struct {
		name      string
		n         int32
		astronaut [][]int32
		want      int64
	}{
		{
			name: "two countries",
			n:    5,
			astronaut: [][]int32{
				{0, 1},
				{2, 3},
				{0, 4},
			},
			want: 6,
		},
		{
			name: "one connected pair and two individuals",
			n:    4,
			astronaut: [][]int32{
				{0, 2},
			},
			want: 5,
		},
		{
			name:      "all astronauts from different countries",
			n:         4,
			astronaut: [][]int32{},
			want:      6,
		},
		{
			name: "all astronauts from the same country",
			n:    4,
			astronaut: [][]int32{
				{0, 1},
				{1, 2},
				{2, 3},
			},
			want: 0,
		},
		{
			name:      "one astronaut",
			n:         1,
			astronaut: [][]int32{},
			want:      0,
		},
		{
			name:      "no astronauts",
			n:         0,
			astronaut: [][]int32{},
			want:      0,
		},
		{
			name: "three countries with different sizes",
			n:    7,
			astronaut: [][]int32{
				{0, 1},
				{1, 2},
				{3, 4},
			},
			// Countries:
			// {0, 1, 2} size 3
			// {3, 4}    size 2
			// {5}       size 1
			// {6}       size 1
			//
			// Valid pairs:
			// 3*2 + 3*1 + 3*1 + 2*1 + 2*1 + 1*1
			// = 17
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := journeyToMoon(tt.n, tt.astronaut)

			if got != tt.want {
				t.Errorf(
					"journeyToMoon(%d, %v) = %d, want %d",
					tt.n,
					tt.astronaut,
					got,
					tt.want,
				)
			}
		})
	}
}
