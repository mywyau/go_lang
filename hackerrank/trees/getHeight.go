package trees

func getHeight(root *TreeNode) int {

	if root == nil {
		return -1
	}

	leftHeight := getHeight(root.left)
	rightHeight := getHeight(root.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}


