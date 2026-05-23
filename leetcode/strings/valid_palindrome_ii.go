package strings

func ValidPalindromeII(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return isPalindromeRange(s, left+1, right) ||
				isPalindromeRange(s, left, right-1)
		}

		left++
		right--
	}

	return true
}

func isPalindromeRange(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}