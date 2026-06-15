package trees

import "fmt"

type QueueItem struct {
	node *TreeNode
	hd   int
}

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

func topView(root *TreeNode) {

	if root == nil {
		return
	}

	topView := map[int]int{}

	minHD := 0
	maxHD := 0

	queue := []QueueItem{
		{node: root, hd: 0},
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if _, exists := topView[item.hd]; !exists {
			topView[item.hd] = item.node.data
		}

		if item.hd < minHD {
			minHD = item.hd
		}

		if item.hd > maxHD {
			maxHD = item.hd
		}

		if item.node.left != nil {
			queue = append(queue, QueueItem{
				node: item.node.left,
				hd:   item.hd - 1,
			})
		}

		if item.node.right != nil {
			queue = append(queue, QueueItem{
				node: item.node.right,
				hd:   item.hd + 1,
			})
		}
	}

	for hd := minHD; hd <= maxHD; hd++ {
		fmt.Printf("%d ", topView[hd])
	}
}
