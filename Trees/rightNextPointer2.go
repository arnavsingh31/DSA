package trees

/*
	LC #117
	TC---> O(n)
	SC---> O(k)
*/
func Connect2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}

	appendToQueue := func(node *TreeNode) {
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	for len(queue) > 0 {
		prev := queue[0]
		queue = queue[1:]

		levelLength := len(queue)
		appendToQueue(prev)

		for i := 0; i < levelLength; i++ {
			pop := queue[0]
			queue = queue[1:]
			appendToQueue(pop)

			prev.Next = pop
			prev = pop
		}

	}
	return root
}
