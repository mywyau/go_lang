package trees

import (
	"testing"
)

func TestLevelOrderSampleTree(t *testing.T) {
	/*
		        1
		         \
		          2
		           \
		            5
		           / \
		          3   6
		           \
		            4

			Level order:
			1 2 5 3 6 4
	*/

	root := &TreeNode{data: 1}
	root.right = &TreeNode{data: 2}
	root.right.right = &TreeNode{data: 5}
	root.right.right.left = &TreeNode{data: 3}
	root.right.right.right = &TreeNode{data: 6}
	root.right.right.left.right = &TreeNode{data: 4}

	got := captureOutput(func() {
		levelOrder(root)
	})

	want := "1 2 5 3 6 4 "

	if got != want {
		t.Errorf("levelOrder() = %q, want %q", got, want)
	}
}

func TestLevelOrderEmptyTree(t *testing.T) {
	got := captureOutput(func() {
		levelOrder(nil)
	})

	want := ""

	if got != want {
		t.Errorf("levelOrder(nil) = %q, want %q", got, want)
	}
}

func TestLevelOrderSingleNode(t *testing.T) {
	root := &TreeNode{data: 10}

	got := captureOutput(func() {
		levelOrder(root)
	})

	want := "10 "

	if got != want {
		t.Errorf("levelOrder(single node) = %q, want %q", got, want)
	}
}

func TestLevelOrderBalancedTree(t *testing.T) {
	/*
		         1
		       /   \
		      2     3
		     / \   / \
		    4   5 6   7

			Level order:
			1 2 3 4 5 6 7
	*/

	root := &TreeNode{data: 1}
	root.left = &TreeNode{data: 2}
	root.right = &TreeNode{data: 3}
	root.left.left = &TreeNode{data: 4}
	root.left.right = &TreeNode{data: 5}
	root.right.left = &TreeNode{data: 6}
	root.right.right = &TreeNode{data: 7}

	got := captureOutput(func() {
		levelOrder(root)
	})

	want := "1 2 3 4 5 6 7 "

	if got != want {
		t.Errorf("levelOrder() = %q, want %q", got, want)
	}
}
