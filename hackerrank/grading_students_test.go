package hackerrank

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	tests := []struct {
		name   string
		grades []int32
		want   []int32
	}{
		{
			name:   "sample case",
			grades: []int32{73, 67, 38, 33},
			want:   []int32{75, 67, 40, 33},
		},
		{
			name:   "grade below 38 does not round",
			grades: []int32{37},
			want:   []int32{37},
		},
		{
			name:   "grade exactly 38 can round",
			grades: []int32{38},
			want:   []int32{40},
		},
		{
			name:   "difference of 1 rounds up",
			grades: []int32{84},
			want:   []int32{85},
		},
		{
			name:   "difference of 2 rounds up",
			grades: []int32{78},
			want:   []int32{80},
		},
		{
			name:   "difference of 3 does not round",
			grades: []int32{57},
			want:   []int32{57},
		},
		{
			name:   "already multiple of 5 stays same",
			grades: []int32{40, 75, 100},
			want:   []int32{40, 75, 100},
		},
		{
			name:   "mixed grades",
			grades: []int32{29, 38, 39, 40, 41, 42, 43, 44, 45},
			want:   []int32{29, 40, 40, 40, 41, 42, 45, 45, 45},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gradingStudents(tt.grades)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gradingStudents(%v) = %v, want %v", tt.grades, got, tt.want)
			}
		})
	}
}