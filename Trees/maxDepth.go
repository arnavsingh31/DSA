package main

/*
	LC #104
	TC---> O(n)
	SC---> O(n)
*/
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftNodes := maxDepth(root.Left) + 1
	rightNodes := maxDepth(root.Right) + 1

	return maxNodes(leftNodes, rightNodes)
}

func maxNodes(a, b int) int {
	if a > b {
		return a
	}
	return b
}
