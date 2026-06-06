package trees

import "fmt"

// Insert a value into the binary search tree.
func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data}
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
	1. If the current node is nil, stop.
	2. Print the current node.
	3. Visit the left child.
	4. Visit the right child.
*/

func preOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func main() {
	var n int
	fmt.Scan(&n)

	var root *Node

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		root = insert(root, value)
	}

	preOrder(root)
}
