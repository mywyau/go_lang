package trees

import "fmt"

/*
Top View of a Binary Tree

Idea:
- Imagine looking at the tree from directly above.
- Each node has a horizontal distance (hd) from the root.
- Root starts at hd = 0.
- Going left decreases hd by 1.
- Going right increases hd by 1.

Example:

        1
         \
          2
           \
            5
           / \
          3   6
           \
            4

Horizontal distances:
- 1 is hd 0
- 2 is hd 1
- 5 is hd 2
- 3 is hd 1
- 6 is hd 3
- 4 is hd 2

Top view means:
- For each horizontal distance, keep the first node we see.
- We use a queue to do level-order traversal.
- This ensures we see higher nodes before lower nodes.

Why BFS / queue?
- Top view depends on which node appears first from the top.
- BFS visits the tree level by level from top to bottom.

Result:
1 2 5 6
*/

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