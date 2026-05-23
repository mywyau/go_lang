package strings

func CustomSortString(order string, s string) string {
	counts := map[byte]int{}

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	result := []byte{}

	for i := 0; i < len(order); i++ {
		ch := order[i]

		for counts[ch] > 0 {
			result = append(result, ch)
			counts[ch]--
		}
	}

	for ch, count := range counts {
		for count > 0 {
			result = append(result, ch)
			count--
		}
	}

	return string(result)
}