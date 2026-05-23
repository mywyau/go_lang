package strings

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example one",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "all same characters",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "mixed repeating characters",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "duplicate outside current window",
			s:    "abba",
			want: 2,
		},
		{
			name: "no repeats",
			s:    "abcdef",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.s)

			if got != tt.want {
				t.Fatalf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
