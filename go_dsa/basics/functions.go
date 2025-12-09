package basics

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add_shorthand(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("div by zero")
	}
	return a / b, nil
}
