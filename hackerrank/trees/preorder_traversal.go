package trees

import "fmt"

// Insert a value into the binary search tree.

func insert(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{data: data}
	}

	if data <= root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

/*
	Preorder means:

	Root -> Left -> Right

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Print the current TreeNode.
	3. Visit the left child.
	4. Visit the right child.
*/

// solution is recursive

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}
