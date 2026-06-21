package warmup

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {

	hourText := s[:2]
	rest := s[2:8]
	period := s[8:]

	hour, _ := strconv.Atoi(hourText)

	if period == "AM" {
		if hour == 12 {
			hour = 0
		}
	} else {
		if hour != 12 {
			hour += 12
		}
	}

	return fmt.Sprintf("%02d%s", hour, rest)
}
