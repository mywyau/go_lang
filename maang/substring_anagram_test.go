package maang

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  int32
	}{
		{
			name:  "two changes required",
			given: "abccde",
			want:  2,
		},
		{
			name:  "already anagrams",
			given: "aaabbb",
			want:  3,
		},
		{
			name:  "same halves",
			given: "abcabc",
			want:  0,
		},
		{
			name:  "odd length",
			given: "abc",
			want:  -1,
		},
		{
			name:  "one change required",
			given: "xaxbbbxx",
			want:  1,
		},
		{
			name:  "empty string",
			given: "",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubstringAnagram(tt.given)

			if got != tt.want {
				t.Errorf(
					"anagram(%q) = %d, want %d",
					tt.given,
					got,
					tt.want,
				)
			}
		})
	}
}
