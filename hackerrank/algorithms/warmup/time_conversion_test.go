package warmup

import "testing"

func TestTimeConversion(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "sample PM time",
			given: "07:05:45PM",
			want:  "19:05:45",
		},
		{
			name:  "12 PM stays as 12",
			given: "12:00:01PM",
			want:  "12:00:01",
		},
		{
			name:  "12 AM becomes 00",
			given: "12:00:01AM",
			want:  "00:00:01",
		},
		{
			name:  "normal AM time stays the same",
			given: "06:40:03AM",
			want:  "06:40:03",
		},
		{
			name:  "11 PM becomes 23",
			given: "11:59:59PM",
			want:  "23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timeConversion(tt.given)

			if got != tt.want {
				t.Errorf("timeConversion(%q) = %q, want %q", tt.given, got, tt.want)
			}
		})
	}
}