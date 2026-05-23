package strings

// Given two strings, return true if they contain the same letters with the same counts.

// s = "anagram"
// t = "nagaram"

// true

// s = "rat"
// t = "car"

// false

// we solve this using a frequency map

// func IsAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	counts := map[rune]int{}

// 	for _, ch := range s {
// 		counts[ch]++
// 	}

// 	for _, ch := range t {
// 		counts[ch]--

// 		if counts[ch] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := map[rune]int{}

	for _, ch := range s {
		counts[ch]++
	}

	for _, ch := range s {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}
