package trees

import "math"

/*
	LC #222
	TC---> < O(n), where n is the number of nodes, since we are not traversing every node in tree.
	SC---> < O(n)
*/
func CountNodes(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := 0
	rightHeight := 0
	node1 := root
	for node1 != nil {
		leftHeight++
		node1 = node1.Left
	}
	node2 := root
	for node2 != nil {
		rightHeight++
		node2 = node2.Right
	}

	if rightHeight == leftHeight {
		return int(math.Pow(2, float64(leftHeight))) - 1
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)

}
