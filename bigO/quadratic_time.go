package bigo

import "fmt"

// O(n^2)
// often nested loops but not necessarily
// if nums has 5 items, 5 x 5 = 25 pairs
// if nums has 10 items then it's 100 x 100 = 10,000 pairs

func printPairs(nums []int) {
	for _, a := range nums {
		for _, b := range nums {
			fmt.Println(a, b)
		}
	}
}
