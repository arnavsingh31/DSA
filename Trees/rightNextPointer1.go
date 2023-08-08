package main

/*
	LC #116
	TC---> O(n)
	SC--->O(n)
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func connect(root *TreeNode) *TreeNode {
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
