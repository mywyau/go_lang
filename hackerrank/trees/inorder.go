package trees

import "fmt"

/*
	inorder means:

	Root -> Left -> Right

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Print the current TreeNode.
	3. Visit the left child.
	4. Visit the right child.
*/

// solution is recursive

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.left)
	fmt.Printf("%d ", root.data)
	inorder(root.right)
}
