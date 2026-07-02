package strings

import (
	"strconv"
)

/*
	Time Conversion

	Problem:
	Given a time in 12-hour AM/PM format, convert it to 24-hour
	military time.

	Example:
	"07:05:45PM" -> "19:05:45"

	Input format:
	The input string always looks like this:

	"hh:mm:ssAM"
	or
	"hh:mm:ssPM"

	Important rules:
	- 12:00:00AM is midnight, so it becomes 00:00:00
	- 12:00:00PM is noon, so it stays 12:00:00
	- AM times stay the same, except for 12AM
	- PM times add 12 to the hour, except for 12PM

	String slicing:
	s[:2]   gets the hour
	s[2:8]  gets the minutes and seconds
	s[8:]   gets AM or PM

	Example:
	s := "07:05:45PM"

	s[:2]  == "07"
	s[2:8] == ":05:45"
	s[8:]  == "PM"
*/

func timeConversion(s string) string {
	hour := s[:2]
	rest := s[2:8]
	period := s[8:]

	if period == "AM" {
		if hour == "12" {
			return "00" + rest
		}

		return hour + rest
	}

	// PM case
	if hour == "12" {
		return "12" + rest
	}

	hourInt, _ := strconv.Atoi(hour)
	hourInt += 12

	return strconv.Itoa(hourInt) + rest
}

// func timeConversion(s string) string {
// 	hour := s[:2]
// 	rest := s[2:8]
// 	period := s[8:]

// 	if period == "AM" {
// 		if hour == "12" {
// 			return "00" + rest
// 		}
// 		return hour + rest
// 	}

// 	if hour == "12" {
// 		return "12" + rest
// 	}

// 	hourInt, _ := strconv.Atoi(hour)
// 	hourInt += 12
// 	return strconv.Itoa(hourInt) + rest
// }