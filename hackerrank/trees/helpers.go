package trees

// Insert a value into the binary search tree.

func Insert(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{data: data}
	}

	if data <= root.data {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}

	return root
}
