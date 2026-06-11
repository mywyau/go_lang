package trees

import "fmt"

/*
	Inorder traversal means:

	Left -> Root -> Right

	Simple idea:
	- Visit everything on the left first.
	- Then print the current node.
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

	Inorder traversal:

	Left -> Root -> Right

	Output:

	1 2 3 4 5 6

	Why?

	- Node 1 has no left child, so print 1.
	- Go right to 2.
	- Node 2 has no left child, so print 2.
	- Go right to 5.
	- Before printing 5, visit its left side.
	- Visit 3.
	- 3 has no left child, so print 3.
	- Go right to 4, print 4.
	- Now go back and print 5.
	- Then visit 5's right child and print 6.

	Steps:
	1. If the current TreeNode is nil, stop.
	2. Visit the left child.
	3. Print the current TreeNode.
	4. Visit the right child.

	Important:
	For a Binary Search Tree, inorder traversal prints the values in sorted order.
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

// func main() {
//     var n int
//     fmt.Scan(&n)

//     var root *Node

//     for i := 0; i < n; i++ {
//         var value int
//         fmt.Scan(&value)

//         root = insert(root, value)
//     }

//     inOrder(root)
// }


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