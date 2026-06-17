package trees

import "fmt"

/*
Level Order Traversal

Idea:
- Visit the tree level by level, from left to right.
- Use a queue because we want to process nodes in the same order we discover them.
- Start by putting the root node into the queue.
- Repeatedly:
  - take the first node from the queue
  - print it
  - add its left child
  - add its right child

Why a queue?
- A queue is FIFO: first in, first out.
- This means nodes from earlier levels are processed before nodes from deeper levels.

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

Output:
	1 2 5 3 6 4
*/

func levelOrder(root *TreeNode) {
	
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", current.data)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}
