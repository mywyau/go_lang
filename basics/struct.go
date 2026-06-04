package basics

import "fmt"

// Person is a struct.
//
// A struct lets us group related data together.
// In this example, a Person has:
// - Name
// - Age
//
// This is similar to a "class with fields" in other languages,
// but Go does not have classes. Go uses structs + methods instead.
type Person struct {
	Name string
	Age  int
}

// NewPerson is a constructor-style function.
//
// Go does not have constructors like Java, C#, or TypeScript.
// Instead, it is common to create a normal function that returns a struct.
//
// This function creates and returns a new Person.
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// Greet is a method on Person.
//
// The part before the function name is called a receiver:
//
//	func (p Person) Greet()
//
// This means Greet belongs to the Person type.
//
// This uses a value receiver:
//
//	p Person
//
// A value receiver gets a copy of the struct.
// This is good when you only need to read data.
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// HaveBirthday is also a method on Person.
//
// This uses a pointer receiver:
//
//	p *Person
//
// A pointer receiver lets us modify the original struct.
//
// We use this because we want to update the person's Age.
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Println(p.Name, "is now", p.Age)
}

// Rename is another pointer receiver method.
//
// Because this method changes Name,
// it needs a pointer receiver.
func (p *Person) Rename(newName string) {
	p.Name = newName
}

// StructExample shows how to create and use structs.
//
// In Go, we usually would not call this ClassExample,
// because Go does not have classes.
func StructExample() {
	// Create a new Person using our constructor-style function.
	p := NewPerson("John", 30)

	// Call a value receiver method.
	// This reads from the struct.
	p.Greet()

	// Call a pointer receiver method.
	// This modifies the original struct.
	p.HaveBirthday()

	// Rename also modifies the original struct.
	p.Rename("Michael")

	// Greet again to show the updated name.
	p.Greet()

	fmt.Println("Final person:", p)
}