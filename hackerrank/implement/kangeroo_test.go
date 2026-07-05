package implement

import "testing"

func TestKangaroo(t *testing.T) {
	tests := []struct {
		name string
		x1   int32
		v1   int32
		x2   int32
		v2   int32
		want string
	}{
		{
			name: "kangaroos meet after 4 jumps",
			x1:   0,
			v1:   3,
			x2:   4,
			v2:   2,
			want: "YES",
		},
		{
			name: "kangaroo behind is slower",
			x1:   0,
			v1:   2,
			x2:   5,
			v2:   3,
			want: "NO",
		},
		{
			name: "kangaroo behind is faster but does not land exactly",
			x1:   0,
			v1:   4,
			x2:   5,
			v2:   2,
			want: "NO",
		},
		{
			name: "kangaroos start apart with same speed",
			x1:   0,
			v1:   3,
			x2:   4,
			v2:   3,
			want: "NO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kangaroo(tt.x1, tt.v1, tt.x2, tt.v2)

			if got != tt.want {
				t.Errorf(
					"kangaroo(%d, %d, %d, %d) = %q, want %q",
					tt.x1,
					tt.v1,
					tt.x2,
					tt.v2,
					got,
					tt.want,
				)
			}
		})
	}
}
