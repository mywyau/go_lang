package strings

func GroupStrings(input []string) [][]string {
	groups := map[string][]string{}

	for _, word := range input {
		key := shiftedKey(word)
		groups[key] = append(groups[key], word)
	}

	result := [][]string{}

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func shiftedKey(word string) string {
	if len(word) == 1 {
		return "single"
	}

	key := []byte{}

	for i := 1; i < len(word); i++ {
		diff := int(word[i] - word[i-1])

		if diff < 0 {
			diff += 26
		}

		key = append(key, byte(diff)+'a')
		key = append(key, '#')
	}

	return string(key)
}