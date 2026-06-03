package basics

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	originalStdout := os.Stdout

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = writer

	fn()

	writer.Close()
	os.Stdout = originalStdout

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		t.Fatalf("failed to read output: %v", err)
	}

	return buffer.String()
}

func TestDogSpeak(t *testing.T) {
	dog := Dog{Name: "Rex"}

	got := dog.Speak()
	want := "Rex says woof!"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestAddTen(t *testing.T) {
	num := 10

	addTen(&num)

	want := 20
	if num != want {
		t.Fatalf("expected %d, got %d", want, num)
	}
}

func TestDivideSuccess(t *testing.T) {
	got, err := divide(10, 2)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := 5
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestDivideByZero(t *testing.T) {
	got, err := divide(10, 0)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if got != 0 {
		t.Fatalf("expected result to be 0, got %d", got)
	}

	wantMessage := "cannot divide by zero"
	if err.Error() != wantMessage {
		t.Fatalf("expected error message %q, got %q", wantMessage, err.Error())
	}
}

func TestUserStruct(t *testing.T) {
	user := User{
		Name: "Michael",
		Age:  30,
	}

	if user.Name != "Michael" {
		t.Fatalf("expected name Michael, got %s", user.Name)
	}

	if user.Age != 30 {
		t.Fatalf("expected age 30, got %d", user.Age)
	}
}

func TestListNodeStruct(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}

	node1.Next = node2

	if node1.Val != 1 {
		t.Fatalf("expected first node value 1, got %d", node1.Val)
	}

	if node1.Next.Val != 2 {
		t.Fatalf("expected next node value 2, got %d", node1.Next.Val)
	}
}

func TestTreeNodeStruct(t *testing.T) {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}

	if root.Val != 10 {
		t.Fatalf("expected root value 10, got %d", root.Val)
	}

	if root.Left.Val != 5 {
		t.Fatalf("expected left value 5, got %d", root.Left.Val)
	}

	if root.Right.Val != 15 {
		t.Fatalf("expected right value 15, got %d", root.Right.Val)
	}
}

func TestExampleFunctionsPrintExpectedOutput(t *testing.T) {
	tests := []struct {
		name     string
		fn       func()
		contains string
	}{
		{
			name:     "integers example",
			fn:       integersExample,
			contains: "--- Integers ---",
		},
		{
			name:     "floats example",
			fn:       floatsExample,
			contains: "--- Floats ---",
		},
		{
			name:     "strings example",
			fn:       stringsExample,
			contains: "--- Strings ---",
		},
		{
			name:     "booleans example",
			fn:       booleansExample,
			contains: "--- Booleans ---",
		},
		{
			name:     "arrays example",
			fn:       arraysExample,
			contains: "--- Arrays ---",
		},
		{
			name:     "slices example",
			fn:       slicesExample,
			contains: "--- Slices ---",
		},
		{
			name:     "maps example",
			fn:       mapsExample,
			contains: "--- Maps ---",
		},
		{
			name:     "structs example",
			fn:       structsExample,
			contains: "--- Structs ---",
		},
		{
			name:     "pointers example",
			fn:       pointersExample,
			contains: "--- Pointers ---",
		},
		{
			name:     "runes and bytes example",
			fn:       runesAndBytesExample,
			contains: "--- Runes and Bytes ---",
		},
		{
			name:     "interfaces example",
			fn:       interfacesExample,
			contains: "--- Interfaces ---",
		},
		{
			name:     "errors example",
			fn:       errorsExample,
			contains: "--- Errors ---",
		},
		{
			name:     "zero values example",
			fn:       zeroValuesExample,
			contains: "--- Zero Values ---",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(t, tt.fn)

			if !strings.Contains(output, tt.contains) {
				t.Fatalf("expected output to contain %q, got:\n%s", tt.contains, output)
			}
		})
	}
}

func TestMainRunsAllExamples(t *testing.T) {
	output := captureOutput(t, main)

	expectedSections := []string{
		"--- Integers ---",
		"--- Floats ---",
		"--- Strings ---",
		"--- Booleans ---",
		"--- Arrays ---",
		"--- Slices ---",
		"--- Maps ---",
		"--- Structs ---",
		"--- Pointers ---",
		"--- Runes and Bytes ---",
		"--- Interfaces ---",
		"--- Errors ---",
		"--- Zero Values ---",
	}

	for _, section := range expectedSections {
		if !strings.Contains(output, section) {
			t.Fatalf("expected output to contain %q, got:\n%s", section, output)
		}
	}
}