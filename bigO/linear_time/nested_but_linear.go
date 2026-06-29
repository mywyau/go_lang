package bigo

import "fmt"

func printTenThingsForEach(nums []int) {
	for _, num := range nums {
		for i := 0; i < 10; i++ {
			fmt.Println(num, i)
		}
	}
}
