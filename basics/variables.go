package basics

var X int = 10
var Name string = "John"
var Alive bool = true

// common var types
var name string = "Michael"
var age int = 30
var price float64 = 19.99
var active bool = true

// short form vars in functions not allowed in package level
func shortForm() {
	name := "Michael"
	age := 30
	isDeveloper := true

	println(name)
	println(age)
	println(isDeveloper)
	
}
