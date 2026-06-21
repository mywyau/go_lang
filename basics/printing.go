package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Go Printing Notes

The main package for printing in Go is:

	import "fmt"

The three most common printing functions are:

	fmt.Print()    // prints values as given, no automatic newline
	fmt.Println()  // prints values and adds a newline
	fmt.Printf()   // formatted printing using verbs like %s, %d, %v

Common formatting verbs:

	%s   // string
	%d   // integer
	%f   // float
	%.2f // float with 2 decimal places
	%v   // default format, useful when unsure
	%+v  // struct with field names
	%#v  // Go-syntax representation
	%c   // character
	%q   // quoted string or character
	%t   // boolean
	%T   // type of value

Mental model:

	Use Print when you want full control over spaces/newlines.
	Use Println for simple line-by-line output.
	Use Printf when you want formatted output.
	Use Sprintf when you want to create a formatted string without printing it.
	Use Fprintln/Fprintf with bufio.NewWriter(os.Stdout) for HackerRank-style output.
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	printExample()
	printlnExample()
	printfExample()
	formatVerbExamples()
	structPrintingExample()
	slicePrintingExample()
	mapPrintingExample()
	stringCharacterPrintingExample()
	sprintfExample()
	errorfExample()
	bufferedWriterExample()
	printLoopExamples()
}

func printExample() {
	fmt.Println("---- fmt.Print example ----")

	// fmt.Print prints exactly what you give it.
	// It does not automatically add spaces or new lines.

	fmt.Print("Hello")
	fmt.Print("World")

	// Output:
	// HelloWorld

	fmt.Print("\n")

	// If you want spaces, add them yourself.
	fmt.Print("Hello ")
	fmt.Print("World")

	fmt.Print("\n\n")
}

func printlnExample() {
	fmt.Println("---- fmt.Println example ----")

	// fmt.Println prints values and adds a newline at the end.
	fmt.Println("Hello")
	fmt.Println("World")

	// If you pass multiple values, Println adds spaces between them.
	name := "Michael"
	age := 30

	fmt.Println(name, age)
	fmt.Println("Name:", name, "Age:", age)

	fmt.Println()
}

func printfExample() {
	fmt.Println("---- fmt.Printf example ----")

	name := "Michael"
	age := 30

	// %s means "place a string here"
	// %d means "place an integer here"
	// \n means newline

	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	// Printf does not automatically add a newline.
	// So this will print on one line:
	fmt.Printf("Hello")
	fmt.Printf("World")

	fmt.Print("\n")

	// Add \n if you want a new line.
	fmt.Printf("Hello\n")
	fmt.Printf("World\n")

	fmt.Println()
}

func formatVerbExamples() {
	fmt.Println("---- formatting verbs ----")

	name := "Michael"
	age := 30
	price := 12.5
	active := true
	letter := 'A'

	fmt.Printf("string with %%s: %s\n", name)
	fmt.Printf("integer with %%d: %d\n", age)
	fmt.Printf("float with %%f: %f\n", price)
	fmt.Printf("float with 2 decimal places using %%.2f: %.2f\n", price)
	fmt.Printf("default format with %%v: %v\n", active)
	fmt.Printf("character with %%c: %c\n", letter)
	fmt.Printf("quoted string with %%q: %q\n", name)
	fmt.Printf("boolean with %%t: %t\n", active)
	fmt.Printf("type with %%T: %T\n", price)

	fmt.Println()
}

func structPrintingExample() {
	fmt.Println("---- printing structs ----")

	person := Person{
		Name: "Michael",
		Age:  30,
	}

	// Println and %v print the values without field names.
	fmt.Println(person)
	fmt.Printf("%v\n", person)

	// %+v includes field names.
	// This is very useful for debugging.
	fmt.Printf("%+v\n", person)

	// %#v prints the Go-syntax representation.
	// This is also useful when debugging.
	fmt.Printf("%#v\n", person)

	fmt.Println()
}

func slicePrintingExample() {
	fmt.Println("---- printing slices ----")

	numbers := []int{10, 20, 30}
	names := []string{"Michael", "Alice", "Bob"}

	fmt.Println(numbers)
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	fmt.Println(names)
	fmt.Printf("%v\n", names)
	fmt.Printf("%#v\n", names)

	fmt.Println()
}

func mapPrintingExample() {
	fmt.Println("---- printing maps ----")

	scores := map[string]int{
		"alice":   10,
		"bob":     20,
		"michael": 30,
	}

	fmt.Println(scores)
	fmt.Printf("%v\n", scores)
	fmt.Printf("%#v\n", scores)

	// Important:
	// Map order is not guaranteed in Go.
	// So the printed order may change between runs.

	fmt.Println()
}

func stringCharacterPrintingExample() {
	fmt.Println("---- printing characters from strings ----")

	s := "hello"

	// This prints the byte value of the first character.
	fmt.Println(s[0]) // 104

	// Use %c to print it as a character.
	fmt.Printf("%c\n", s[0]) // h

	// Or convert it to a string.
	fmt.Println(string(s[0])) // h

	// Useful example:
	fmt.Printf("first byte as number: %d\n", s[0])
	fmt.Printf("first byte as character: %c\n", s[0])
	fmt.Printf("first byte as string: %s\n", string(s[0]))

	fmt.Println()
}

func sprintfExample() {
	fmt.Println("---- fmt.Sprintf example ----")

	name := "Michael"
	score := 95

	// fmt.Sprintf formats a string but does not print it.
	// It returns the formatted string.
	message := fmt.Sprintf("%s scored %d points", name, score)

	fmt.Println(message)

	// This is useful when you want to build a string and return it.
	anotherMessage := createScoreMessage("Alice", 88)
	fmt.Println(anotherMessage)

	fmt.Println()
}

func createScoreMessage(name string, score int) string {
	return fmt.Sprintf("%s scored %d points", name, score)
}

func errorfExample() {
	fmt.Println("---- fmt.Errorf example ----")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	fmt.Println()
}

func divide(a, b int) (int, error) {
	if b == 0 {
		// fmt.Errorf creates a formatted error.
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}

	return a / b, nil
}

func bufferedWriterExample() {
	fmt.Println("---- buffered writer example ----")

	/*
	You often see this style in HackerRank:

		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()

		fmt.Fprintln(writer, result)

	Why?

	Because writing through a buffer can be faster than printing directly
	many times.

	These are the writer versions:

		fmt.Fprint(writer, ...)
		fmt.Fprintln(writer, ...)
		fmt.Fprintf(writer, ...)

	They are like:

		fmt.Print(...)
		fmt.Println(...)
		fmt.Printf(...)

	But instead of writing directly to stdout, they write to the given writer.
	*/

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, "Hello from buffered writer")
	fmt.Fprintf(writer, "Name: %s, Age: %d\n", "Michael", 30)

	// Because this writes to a buffered writer, Flush sends the output out.
	// We used defer writer.Flush(), so it flushes at the end of main.
}

func printLoopExamples() {
	fmt.Println("---- printing in loops ----")

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Each number on a separate line:")

	for _, number := range numbers {
		fmt.Println(number)
	}

	fmt.Println("Numbers on one line with spaces:")

	for _, number := range numbers {
		fmt.Print(number, " ")
	}

	fmt.Println()

	fmt.Println("Numbers on one line without trailing space:")

	for i, number := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(number)
	}

	fmt.Println()
	fmt.Println()
}