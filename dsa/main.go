package main

import (
	"fmt"
	"go_lang/dsa/basics"
	"go_lang/dsa/searching"
)

func main() {

	arr := []int{1, 3, 5, 7, 9, 11}

	fmt.Println(basics.X, basics.Name, basics.Alive)
	fmt.Println(searching.BinarySearch(arr, 7))
}
