package trees

import "fmt"

/*
	Postorder traversal means:

	Left -> Right -> Root

	Simple idea:
	- Visit everything on the left first.
	- Then visit everything on the right.
	- Print the current node last.

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

	Postorder traversal:

	Left -> Right -> Root

	Output:

	4 3 6 5 2 1

	Why?

	- Start at 1.
	- 1 has no left child.
	- Go right to 2.
	- 2 has no left child.
	- Go right to 5.
	- Before printing 5, visit its left side.
	- Go to 3.
	- 3 has no left child.
	- Go right to 4.
	- 4 has no children, so print 4.
	- Go back to 3 and print 3.
	- Now visit 5's right child, 6.
	- 6 has no children, so print 6.
	- Now print 5.
	- Now print 2.
	- Finally print 1.

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Visit the left child.
	3. Visit the right child.
	4. Print the current TreeNode.

	Important:
	Postorder is useful when you need to process children before the parent.
	For example, deleting/freeing a tree or calculating values from child nodes.
*/

// solution is recursive

func postorder(root *TreeNode) {
	if root == nil {
		return
	}

	postorder(root.left)
	postorder(root.right)
	fmt.Printf("%d ", root.data)
}