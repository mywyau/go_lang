package strings

func LengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}

	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if previousIndex, found := lastSeen[ch]; found && previousIndex >= left {
			left = previousIndex + 1
		}

		lastSeen[ch] = right

		windowLength := right - left + 1

		if windowLength > best {
			best = windowLength
		}
	}

	return best
}
