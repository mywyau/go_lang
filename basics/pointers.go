package basics

import "fmt"

// Pointers let us store the memory address of a value.
//
// The two most important symbols are:
//
//	& = "address of"
//	* = "value at address" OR "pointer type"
//
// Example:
//
//	x := 10
//	p := &x
//
// x stores the actual value:
//
//	x = 10
//
// p stores the memory address of x:
//
//	p = address of x
//
// So p does not store 10 directly.
// It stores where x lives in memory.

func getMemory() {
	x := 10

	fmt.Println("value of x:", x)

	// &x means:
	// "give me the memory address of x"
	fmt.Println("address of x:", &x)
}

func createPointer() {
	x := 10

	// p stores the address of x.
	//
	// Because x is an int, p has the type *int.
	//
	// *int means:
	// "pointer to an int"
	p := &x

	fmt.Println("x:", x)
	fmt.Println("p:", p)
}

func dereferencePointer() {
	x := 10
	p := &x

	fmt.Println("address stored in p:", p)

	// *p means:
	// "go to the address stored in p and get the value there"
	//
	// This is called dereferencing.
	fmt.Println("value at that address:", *p)
}

func updateThroughPointer() {
	x := 10
	p := &x

	fmt.Println("before:", x)

	// *p = 20 means:
	// "go to the address stored in p and change the value there to 20"
	//
	// Because p points to x, this changes x.
	*p = 20

	fmt.Println("after:", x)
}

func pointerAsFunctionArgument() {
	x := 10

	fmt.Println("before:", x)

	// We pass the address of x into the function.
	// This allows the function to modify the original x.
	increment(&x)

	fmt.Println("after:", x)
}

// n *int means:
//
//	n is a pointer to an int
//
// So n stores an address, not the actual int directly.
func increment(n *int) {
	// *n means:
	// "go to the int that n points to"
	//
	// Then we increase that value by 1.
	*n++
}

func nilPointer() {
	// A pointer can be nil.
	//
	// nil means:
	// "this pointer does not point to anything yet"
	var p *int

	fmt.Println("p:", p)

	// Be careful:
	// This would crash because p points to nothing:
	//
	// fmt.Println(*p)
}