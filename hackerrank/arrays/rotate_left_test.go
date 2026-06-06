package arrays

import (
	"reflect"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	tests := []struct {
		name string
		d    int32
		arr  []int32
		want []int32
	}{
		{
			name: "rotate by 2",
			d:    2,
			arr:  []int32{1, 2, 3, 4, 5},
			want: []int32{3, 4, 5, 1, 2},
		},
		{
			name: "rotate by 4",
			d:    4,
			arr:  []int32{1, 2, 3, 4, 5},
			want: []int32{5, 1, 2, 3, 4},
		},
		{
			name: "rotate by array length",
			d:    5,
			arr:  []int32{1, 2, 3, 4, 5},
			want: []int32{1, 2, 3, 4, 5},
		},
		{
			name: "rotate more than array length",
			d:    7,
			arr:  []int32{1, 2, 3, 4, 5},
			want: []int32{3, 4, 5, 1, 2},
		},
		{
			name: "single item",
			d:    3,
			arr:  []int32{10},
			want: []int32{10},
		},
		{
			name: "empty array",
			d:    3,
			arr:  []int32{},
			want: []int32{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateLeft(tt.d, tt.arr)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}