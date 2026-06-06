package trees

// import (
// 	"bytes"
// 	"io"
// 	"os"
// 	"testing"
// )

// // Helper to capture anything printed to stdout.
// //
// // This lets us test functions like preOrder()
// // because preOrder prints instead of returning a value.
// func captureOutput(fn func()) string {
// 	old := os.Stdout

// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	fn()

// 	w.Close()
// 	os.Stdout = old

// 	var buf bytes.Buffer
// 	io.Copy(&buf, r)

// 	return buf.String()
// }

// func TestPreOrder(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		values []int
// 		want   string
// 	}{
// 		{
// 			name:   "example tree",
// 			values: []int{1, 2, 5, 3, 6, 4},
// 			want:   "1 2 5 3 4 6 ",
// 		},
// 		{
// 			name:   "single node",
// 			values: []int{10},
// 			want:   "10 ",
// 		},
// 		{
// 			name:   "empty tree",
// 			values: []int{},
// 			want:   "",
// 		},
// 		{
// 			name:   "left side tree",
// 			values: []int{5, 4, 3, 2, 1},
// 			want:   "5 4 3 2 1 ",
// 		},
// 		{
// 			name:   "right side tree",
// 			values: []int{1, 2, 3, 4, 5},
// 			want:   "1 2 3 4 5 ",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var root *Node

// 			for _, value := range tt.values {
// 				root = insert(root, value)
// 			}

// 			got := captureOutput(func() {
// 				preOrder(root)
// 			})

// 			if got != tt.want {
// 				t.Fatalf("expected %q, got %q", tt.want, got)
// 			}
// 		})
// 	}
// }