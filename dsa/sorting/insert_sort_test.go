package sorting

import "testing"

func TestInsertionSort(t *testing.T) {

	got := InsertionSort([]int{2, 7, 15, 11, 1, 0, 100, 300, 89, 91, 52})
	want := []int{0, 1, 2, 7, 11, 15, 52, 89, 91, 100, 300}

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}
