package basics

import "fmt"

/*
Go String Basics

A string in Go is immutable. That means you cannot directly edit a character
inside it. Most string operations create a new string.

For HackerRank-style problems, the most important tools are:

- len(s)
- s[i]
- s[start:end]
- for loops
- range loops
*/

func declaringStrings() {
	fmt.Println("--- declaring strings ---")

	name := "Michael"
	message := "hello Go"

	// Raw strings use backticks. They keep line breaks and do not require escaping quotes.
	raw := `This is a raw string.
It can span multiple lines.
"Quotes" do not need escaping here.`

	fmt.Println(name)
	fmt.Println(message)
	fmt.Println(raw)
}

func stringLength() {
	fmt.Println("--- string length ---")

	s := "hello"

	// len gives the number of bytes in the string.
	fmt.Println(len(s)) // 5

	// For normal ASCII text, bytes and characters usually feel the same.
	// For Unicode text, they can be different.
	fmt.Println(len("é")) // often 2 bytes in UTF-8
}

func indexingStrings() {
	fmt.Println("--- indexing strings ---")

	s := "hello"

	// s[0] gives the first byte, not a string.
	fmt.Println(s[0])        // 104, ASCII/UTF-8 byte value for 'h'
	fmt.Printf("%c\n", s[0]) // h

	first := s[0]
	last := s[len(s)-1]

	fmt.Printf("first = %c, last = %c\n", first, last)
}

func loopingOverStrings() {
	fmt.Println("--- looping over strings ---")

	s := "hello"

	// Byte-based loop. Good for ASCII strings.
	for i := 0; i < len(s); i++ {
		fmt.Printf("byte index %d = %c\n", i, s[i])
	}

	// Rune-based loop. Better for Unicode text.
	for index, ch := range s {
		fmt.Printf("range index %d = %c\n", index, ch)
	}
}

func slicingStrings() {
	fmt.Println("--- slicing strings ---")

	s := "abcdef"

	// Syntax: s[start:end]
	// Includes start, excludes end.
	fmt.Println(s[0:2]) // ab
	fmt.Println(s[2:])  // cdef
	fmt.Println(s[:3])  // abc
	fmt.Println(s[1:4]) // bcd
}

func main() {
	declaringStrings()
	stringLength()
	indexingStrings()
	loopingOverStrings()
	slicingStrings()
}