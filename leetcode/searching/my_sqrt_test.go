package searching

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "sqrt 0",
			x:    0,
			want: 0,
		},
		{
			name: "sqrt 1",
			x:    1,
			want: 1,
		},
		{
			name: "sqrt 4",
			x:    4,
			want: 2,
		},
		{
			name: "sqrt 8",
			x:    8,
			want: 2,
		},
		{
			name: "sqrt 9",
			x:    9,
			want: 3,
		},
		{
			name: "sqrt 15",
			x:    15,
			want: 3,
		},
		{
			name: "sqrt 16",
			x:    16,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MySqrt(tt.x)

			if got != tt.want {
				t.Errorf("MySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}