package arrays

import (
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	// Save the original stdout
	originalStdout := os.Stdout

	// Create a pipe so we can capture printed output
	reader, writer, _ := os.Pipe()

	// Replace stdout with our writer
	os.Stdout = writer

	// Run the function
	f()

	// Close writer and restore stdout
	writer.Close()
	os.Stdout = originalStdout

	// Read everything that was printed
	output, _ := io.ReadAll(reader)

	return string(output)
}

func TestArrayExamples(t *testing.T) {
	output := captureOutput(ArrayExamples)

	if !strings.Contains(output, "array: [10 20 30]") {
		t.Fatalf("expected output to contain array, got:\n%s", output)
	}

	if !strings.Contains(output, "after update: [10 99 30]") {
		t.Fatalf("expected output to contain updated array, got:\n%s", output)
	}

	if !strings.Contains(output, "array length: 3") {
		t.Fatalf("expected output to contain array length, got:\n%s", output)
	}
}

func TestSliceExamples(t *testing.T) {
	output := captureOutput(SliceExamples)

	if !strings.Contains(output, "slice: [1 2 3]") {
		t.Fatalf("expected output to contain original slice, got:\n%s", output)
	}

	if !strings.Contains(output, "after append: [1 2 3 4 5]") {
		t.Fatalf("expected output to contain appended slice, got:\n%s", output)
	}

	if !strings.Contains(output, "after update: [100 2 3 4 5]") {
		t.Fatalf("expected output to contain updated slice, got:\n%s", output)
	}
}

func TestSliceCuttingExamples(t *testing.T) {
	output := captureOutput(SliceCuttingExamples)

	if !strings.Contains(output, "original: [10 20 30 40 50]") {
		t.Fatalf("expected output to contain original slice, got:\n%s", output)
	}

	if !strings.Contains(output, "nums[1:4]: [20 30 40]") {
		t.Fatalf("expected output to contain nums[1:4], got:\n%s", output)
	}

	if !strings.Contains(output, "nums[:3]: [10 20 30]") {
		t.Fatalf("expected output to contain nums[:3], got:\n%s", output)
	}

	if !strings.Contains(output, "nums[2:]: [30 40 50]") {
		t.Fatalf("expected output to contain nums[2:], got:\n%s", output)
	}
}

func TestSliceMakeExamples(t *testing.T) {
	output := captureOutput(SliceMakeExamples)

	if !strings.Contains(output, "made slice: [0 0 0]") {
		t.Fatalf("expected output to contain made slice, got:\n%s", output)
	}

	if !strings.Contains(output, "after updates: [10 20 30]") {
		t.Fatalf("expected output to contain updated made slice, got:\n%s", output)
	}

	if !strings.Contains(output, "after append: [100 200]") {
		t.Fatalf("expected output to contain appended scores slice, got:\n%s", output)
	}
}

func TestArrayVsSliceExamples(t *testing.T) {
	output := captureOutput(ArrayVsSliceExamples)

	if !strings.Contains(output, "array: [1 2 3]") {
		t.Fatalf("expected output to contain array, got:\n%s", output)
	}

	if !strings.Contains(output, "slice: [1 2 3]") {
		t.Fatalf("expected output to contain slice, got:\n%s", output)
	}

	if !strings.Contains(output, "slice after append: [1 2 3 4]") {
		t.Fatalf("expected output to contain appended slice, got:\n%s", output)
	}
}