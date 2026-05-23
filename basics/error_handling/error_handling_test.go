package errorhandling

import (
	"os"
	"testing"
)

func TestReadReturnsErrorWhenFileDoesNotExist(t *testing.T) {
	withTempWorkingDirectory(t)

	err := read()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestReadReturnsNilWhenFileExists(t *testing.T) {
	withTempWorkingDirectory(t)

	err := os.WriteFile("file.txt", []byte("hello"), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	got := read()

	if got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func withTempWorkingDirectory(t *testing.T) {
	t.Helper()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	tempDir := t.TempDir()

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to change working directory: %v", err)
	}

	t.Cleanup(func() {
		err := os.Chdir(originalDir)
		if err != nil {
			t.Fatalf("failed to restore working directory: %v", err)
		}
	})
}
