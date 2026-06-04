package linkedlists

import (
	"testing"
)

// func captureOutput(f func()) string {
// 	// Save the real stdout
// 	originalStdout := os.Stdout

// 	// Create a pipe so we can capture anything printed
// 	reader, writer, _ := os.Pipe()

// 	// Temporarily replace stdout
// 	os.Stdout = writer

// 	// Run the function we want to test
// 	f()

// 	// Close writer and restore stdout
// 	writer.Close()
// 	os.Stdout = originalStdout

// 	// Read the captured output
// 	output, _ := io.ReadAll(reader)

// 	return string(output)
// }

func TestReversePrintMultipleNodes(t *testing.T) {
	node1 := &SinglyLinkedListNode{data: 16}
	node2 := &SinglyLinkedListNode{data: 13}

	node1.next = node2

	output := captureOutput(func() {
		reversePrint(node1)
	})

	want := "16\n13\n"

	if output != want {
		t.Fatalf("expected %q, got %q", want, output)
	}
}

func TestReversePrintSingleNode(t *testing.T) {
	node1 := &SinglyLinkedListNode{data: 99}

	output := captureOutput(func() {
		reversePrint(node1)
	})

	want := "99\n"

	if output != want {
		t.Fatalf("expected %q, got %q", want, output)
	}
}

func TestReversePrintEmptyList(t *testing.T) {
	output := captureOutput(func() {
		reversePrint(nil)
	})

	want := ""

	if output != want {
		t.Fatalf("expected %q, got %q", want, output)
	}
}