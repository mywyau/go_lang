package maang

func anagram(s string) int32 {

	if len(s)%2 != 0 {
		return -1
	}

	middle := len(s) / 2

	var counts [26]int

	// Count the letters in the first half.
	for i := 0; i < middle; i++ {
		index := s[i] - 'a'
		counts[index]++
	}

	// Cancel matching letters from the second half.
	for i := middle; i < len(s); i++ {
		index := s[i] - 'a'
		counts[index]--
	}

	var changes int32

	// Positive values represent unmatched letters
	// remaining in the first half.
	for _, count := range counts {
		if count > 0 {
			changes += int32(count)
		}
	}

	return changes
}
