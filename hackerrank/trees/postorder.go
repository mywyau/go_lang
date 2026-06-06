package trees

import "fmt"

/*
	postorder means:

	Left -> Right -> Root

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Visit the left child.
	3. Visit the right child.
	4. Print the current TreeNode.

	Simple idea:
	- Do the children first.
	- Print the current node last.
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