package trees

import "fmt"

/*
	Preorder traversal means:

	Root -> Left -> Right

	Simple idea:
	- Print the current node first.
	- Then visit everything on the left.
	- Then visit everything on the right.

	Example tree:

	        1
	         \
	          2
	           \
	            5
	           / \
	          3   6
	           \
	            4

	Preorder traversal:

	Root -> Left -> Right

	Output:

	1 2 5 3 4 6

	Why?

	- Start at 1, so print 1.
	- 1 has no left child.
	- Go right to 2, so print 2.
	- 2 has no left child.
	- Go right to 5, so print 5.
	- Visit 5's left side.
	- Go to 3, so print 3.
	- 3 has no left child.
	- Go right to 4, so print 4.
	- Go back to 5.
	- Visit 5's right child, 6, so print 6.

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Print the current TreeNode.
	3. Visit the left child.
	4. Visit the right child.

	Important:
	Preorder is useful when you want to process the parent before its children.
	For example, copying a tree or serialising a tree structure.
*/

// solution is recursive

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.data)
	preorder(root.left)
	preorder(root.right)
}