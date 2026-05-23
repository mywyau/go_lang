package strings

func MinRemoveToMakeValid(s string) string {
	stack := []int{}
	remove := map[int]bool{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				remove[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for _, index := range stack {
		remove[index] = true
	}

	result := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if !remove[i] {
			result = append(result, s[i])
		}
	}

	return string(result)
}