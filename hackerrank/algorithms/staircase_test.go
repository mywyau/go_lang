package algorithms

import "testing"

func TestStaircaseOutput(t *testing.T) {
	tests := []struct {
		name string
		n    int32
		want string
	}{
		{
			name: "sample input 6",
			n:    6,
			want: "     #\n" +
				"    ##\n" +
				"   ###\n" +
				"  ####\n" +
				" #####\n" +
				"######\n",
		},
		{
			name: "size 1",
			n:    1,
			want: "#\n",
		},
		{
			name: "size 3",
			n:    3,
			want: "  #\n" +
				" ##\n" +
				"###\n",
		},
		{
			name: "size 0",
			n:    0,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := staircaseOutput(tt.n)

			if got != tt.want {
				t.Errorf("staircaseOutput(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}