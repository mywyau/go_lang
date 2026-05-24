package basics

import "fmt"

func LoopExamples1() {

	// 1. Classic for loop (like C / Java)

	for i := 0; i < 5; i++ {
		fmt.Println("classic:", i)
	}
}

// 2. While-style loop
// (Go has no "while"; you use for)

func LoopExamples2() {

	j := 0
	for j < 5 {
		fmt.Println("while-style:", j)
		j++
	}
}

// 3. Infinite loop (break to exit)
func LoopExamples3() {
	k := 0
	for {
		fmt.Println("infinite:", k)
		k++
		if k == 3 {
			break
		}
	}
}

func LoopExamples4() {
	// 4. Loop over a slice using range

	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println("slice-range:", index, value)
	}
}

// 5. Loop over a map using range
// (order is NOT guaranteed)
func LoopExamples5() {

	ages := map[string]int{
		"alice": 20,
		"bob":   25,
	}

	for key, value := range ages {
		fmt.Println("map-range:", key, value)
	}

	// 6. Loop over a string (runes)

	for i, r := range "Go!" {
		fmt.Println("string-range:", i, string(r))
	}
}
