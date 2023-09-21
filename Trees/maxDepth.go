package trees

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #104
	TC---> O(n)
	SC---> O(n)
*/
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftNodes := MaxDepth(root.Left) + 1
	rightNodes := MaxDepth(root.Right) + 1

	return util.Max(leftNodes, rightNodes)
}
