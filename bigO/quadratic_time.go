package bigo

import "fmt"

func printPairs(nums []int) {
	for _, a := range nums {
		for _, b := range nums {
			fmt.Println(a, b)
		}
	}
}
