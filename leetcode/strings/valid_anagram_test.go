package strings

import "testing"

func TestValidAnagram(t *testing.T) {
	got := ValidAnagram("anagram", "nagaram")
	want := true

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
