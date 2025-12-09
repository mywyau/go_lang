package basics

import "fmt"

// Struct (data)
type Person struct {
    Name string
    Age  int
}

// Constructor function
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}

// Method (value receiver — read only)
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

// Method (pointer receiver — modifies struct)
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Println(p.Name, "is now", p.Age)
}

func ClassExample() {
    p := NewPerson("John", 30)

    p.Greet()
    p.HaveBirthday()
}
