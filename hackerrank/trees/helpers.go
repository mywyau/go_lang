package trees

import (
	"bytes"
	"io"
	"os"
)

// Helper to capture anything printed to stdout.
//
// This lets us test functions like PreOrder()
// because PreOrder prints instead of returning a value.
func captureOutput(fn func()) string {
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

// InsertTreeNode - Insert a value into the binary search tree.

func InsertTreeNode(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{data: data}
	}

	if data <= root.data {
		root.left = InsertTreeNode(root.left, data)
	} else {
		root.right = InsertTreeNode(root.right, data)
	}

	return root
}
