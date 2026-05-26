package basics

import "fmt"

func getMemory() {
	x := 10

	fmt.Println(x)
	fmt.Println(&x) // give me the memory address of x
}

func createPointer() {
	x := 10

	p := &x // p stores the address of x

	fmt.Println("x:", x)
	fmt.Println("p:", p)
}

func dereferencePointer() {
	x := 10
	p := &x

	fmt.Println("address:", p)
	fmt.Println("value:", *p) 		// go to the address stored in p and get the value
}
