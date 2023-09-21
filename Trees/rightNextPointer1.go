package trees

/*
	LC #116
	TC---> O(n)
	SC--->O(k), atmost k elements will be in queue where k is the no of leaf nodes.
*/

func Connect(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}

	addToQueue := func(node *TreeNode) {
		if node != nil && node.Left != nil {
			queue = append(queue, node.Left, node.Right)
		}
	}

	for len(queue) > 0 {
		prev := queue[0]
		queue = queue[1:]

		currLevelLength := len(queue)

		addToQueue(prev)

		for i := 0; i < currLevelLength; i++ {
			pop := queue[0]
			queue = queue[1:]

			prev.Next = pop
			addToQueue(pop)
			prev = pop
		}
	}

	return root
}

/*
**more intuitve BFS same code works for LC 117 as well.
 */
func ConnectTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelLength := len(queue)
		var prev *TreeNode

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			prev = node
		}
	}
	return root
}
