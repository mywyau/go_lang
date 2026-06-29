package bigo

import "fmt"

// this is a nested loop but the inner loop only runs  10 times hence not quadratic since there is constant number of iterations in the inner loop

func printTenThingsForEach(nums []int) {
	for _, num := range nums {
		for i := 0; i < 10; i++ {
			fmt.Println(num, i)
		}
	}
}
