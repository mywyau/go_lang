package basics

import "fmt"

// This file shows common Go types with simple examples.
//
// You can run it with:
//
//   go run common_types_examples.go
//
// Main idea:
// - Go is statically typed.
// - Every variable has a type.
// - Some types are very common in DSA/interview questions:
//   int, string, bool, slices, maps, structs, pointers.

// A struct lets us create our own custom type.
type User struct {
	Name string
	Age  int
}

// Common DSA-style struct: linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Common DSA-style struct: binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// An interface describes behaviour.
// Any type with a Speak() string method satisfies this interface.
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says woof!"
}

func main() {
	integersExample()
	floatsExample()
	stringsExample()
	booleansExample()
	arraysExample()
	slicesExample()
	mapsExample()
	structsExample()
	pointersExample()
	runesAndBytesExample()
	interfacesExample()
	errorsExample()
	zeroValuesExample()
}

func integersExample() {
	fmt.Println("--- Integers ---")

	// int is the most common integer type.
	// Use int most of the time unless you specifically need a sized integer.
	var age int = 30
	score := 100

	fmt.Println("age:", age)
	fmt.Println("score:", score)

	// Signed integers can be negative or positive.
	// The number means how many bits are used.
	var a int8 = 127         // range: -128 to 127
	var b int16 = 32767      // range: -32768 to 32767
	var c int32 = 2147483647 // also used by rune
	var d int64 = 9223372036854775807

	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)

	// Unsigned integers cannot be negative.
	var u uint = 10
	var u8 uint8 = 255 // also used by byte
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Println("uint:", u)
	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)

	// Type conversion is explicit in Go.
	// You cannot automatically add int and int32.
	var x int = 10
	var y int32 = 20

	// total := x + y // This would not compile.
	total := x + int(y)
	fmt.Println("converted total:", total)
}

func floatsExample() {
	fmt.Println("\n--- Floats ---")

	// float64 is the default and most common decimal type.
	var price float64 = 19.99
	pi := 3.14159

	fmt.Println("price:", price)
	fmt.Println("pi:", pi)

	// float32 uses less memory but is less precise.
	var smallerFloat float32 = 1.5
	fmt.Println("float32:", smallerFloat)
}

func stringsExample() {
	fmt.Println("\n--- Strings ---")

	name := "Michael"
	message := "Hello, Go!"

	fmt.Println("name:", name)
	fmt.Println("message:", message)

	fullName := "Michael" + " " + "Yau"
	fmt.Println("full name:", fullName)

	// Indexing a string gives you a byte, not a string.
	word := "hello"
	fmt.Println("word[0] as byte:", word[0])
	fmt.Printf("word[0] as character: %c\n", word[0])
}

func booleansExample() {
	fmt.Println("\n--- Booleans ---")

	isActive := true
	isLoggedIn := false

	fmt.Println("isActive:", isActive)
	fmt.Println("isLoggedIn:", isLoggedIn)

	if isActive {
		fmt.Println("User is active")
	}
}

func arraysExample() {
	fmt.Println("\n--- Arrays ---")

	// Arrays have a fixed size.
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println("numbers:", numbers)
	fmt.Println("first number:", numbers[0])

	// Go can infer the array size using ...
	moreNumbers := [...]int{10, 20, 30}
	fmt.Println("moreNumbers:", moreNumbers)
}

func slicesExample() {
	fmt.Println("\n--- Slices ---")

	// Slices are like dynamic arrays.
	// These are much more common than arrays in Go.
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("nums:", nums)

	nums = append(nums, 6)
	fmt.Println("after append:", nums)

	// Loop over a slice with range.
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}

	// Common DSA example: sum a slice.
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum:", total)
}

func mapsExample() {
	fmt.Println("\n--- Maps ---")

	// A map stores key-value pairs.
	ages := map[string]int{
		"Michael": 30,
		"Alice":   25,
	}

	fmt.Println("ages:", ages)
	fmt.Println("Michael's age:", ages["Michael"])

	// Add a new key/value.
	ages["Bob"] = 40
	fmt.Println("after adding Bob:", ages)

	// Check if a key exists.
	age, exists := ages["Charlie"]
	if exists {
		fmt.Println("Charlie age:", age)
	} else {
		fmt.Println("Charlie not found")
	}

	// Common DSA example: map from number to index.
	seen := map[int]int{}
	seen[10] = 0
	seen[20] = 1
	fmt.Println("seen:", seen)
}

func structsExample() {
	fmt.Println("\n--- Structs ---")

	user := User{
		Name: "Michael",
		Age:  30,
	}

	fmt.Println("user:", user)
	fmt.Println("user name:", user.Name)
	fmt.Println("user age:", user.Age)

	// Linked list example.
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node1.Next = node2

	fmt.Println("linked list first value:", node1.Val)
	fmt.Println("linked list next value:", node1.Next.Val)

	// Binary tree example.
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}

	fmt.Println("tree root:", root.Val)
	fmt.Println("tree left:", root.Left.Val)
	fmt.Println("tree right:", root.Right.Val)
}

func pointersExample() {
	fmt.Println("\n--- Pointers ---")

	num := 10
	pointerToNum := &num

	fmt.Println("num:", num)
	fmt.Println("pointer address:", pointerToNum)
	fmt.Println("value through pointer:", *pointerToNum)

	// Update the original value through the pointer.
	*pointerToNum = 100
	fmt.Println("num after pointer update:", num)

	// Passing pointer to a function.
	addTen(&num)
	fmt.Println("num after addTen:", num)
}

func addTen(value *int) {
	*value += 10
}

func runesAndBytesExample() {
	fmt.Println("\n--- Runes and Bytes ---")

	// byte is an alias for uint8.
	// Good for raw bytes or ASCII characters.
	var b byte = 'a'
	fmt.Println("byte value:", b)
	fmt.Printf("byte as char: %c\n", b)

	// rune is an alias for int32.
	// Good for Unicode characters.
	var r rune = '好'
	fmt.Println("rune value:", r)
	fmt.Printf("rune as char: %c\n", r)

	// range over a string gives runes.
	text := "Go好"
	for index, char := range text {
		fmt.Println("index:", index, "rune:", char)
		fmt.Printf("character: %c\n", char)
	}
}

func interfacesExample() {
	fmt.Println("\n--- Interfaces ---")

	dog := Dog{Name: "Rex"}
	makeItSpeak(dog)
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func errorsExample() {
	fmt.Println("\n--- Errors ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10 / 2 =", result)

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10 / 0 =", result)
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func zeroValuesExample() {
	fmt.Println("\n--- Zero Values ---")

	// If you declare a variable without giving it a value,
	// Go gives it a default zero value.
	var age int
	var price float64
	var name string
	var active bool
	var nums []int
	var user *User
	var ages map[string]int

	fmt.Println("zero int:", age)
	fmt.Println("zero float64:", price)
	fmt.Println("zero string:", name)
	fmt.Println("zero bool:", active)
	fmt.Println("zero slice is nil:", nums == nil)
	fmt.Println("zero pointer is nil:", user == nil)
	fmt.Println("zero map is nil:", ages == nil)
}