package main

import (
	"fmt"	
	"github.com/mywyau/go_lang/go_dsa/searching"
)

func main() {

	arr := []int{1, 3, 5, 7, 9, 11}
    fmt.Println(searching.BinarySearch(arr, 7))
}
