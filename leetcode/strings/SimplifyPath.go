package strings

func SimplifyPath(path string) string {
	stack := []string{}
	i := 0

	for i < len(path) {
		for i < len(path) && path[i] == '/' {
			i++
		}

		start := i

		for i < len(path) && path[i] != '/' {
			i++
		}

		part := path[start:i]

		if part == "" || part == "." {
			continue
		}

		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	result := ""

	for _, part := range stack {
		result += "/" + part
	}

	return result
}
