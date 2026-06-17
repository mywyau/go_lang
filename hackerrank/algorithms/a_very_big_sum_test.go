package algorithms

import "testing"

func TestAVeryBigSum(t *testing.T) {

	arr := []int64{
		1000000001,
		1000000002,
		1000000003,
		1000000004,
		1000000005,
	}

	got := aVeryBigSum(arr)
	want := int64(5000000015)

	if got != want {
		t.Errorf("expected %d , got %d", want, got)
	}
}
