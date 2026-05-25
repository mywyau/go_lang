package binary

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "16 is perfect square",
			num:  16,
			want: true,
		},
		{
			name: "14 is not perfect square",
			num:  14,
			want: false,
		},
		{
			name: "1 is perfect square",
			num:  1,
			want: true,
		},
		{
			name: "25 is perfect square",
			num:  25,
			want: true,
		},
		{
			name: "26 is not perfect square",
			num:  26,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPerfectSquare(tt.num)

			if got != tt.want {
				t.Errorf("IsPerfectSquare(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}