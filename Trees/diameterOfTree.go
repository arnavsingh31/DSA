package main

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #543
	TC--->O(n)
	SC--->O(n)
*/
func DiameterOfTree(root *Node) int {
	ans := 0
	solveDiameter(root, 0, &ans)
	return ans
}

func solveDiameter(root *Node, edge int, ans *int) int {
	if root == nil {
		// returning -1 because we don't want to add it to the path/edge in case where one child doesn't exist
		return -1
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	left := 1 + solveDiameter(root.Left, edge, ans)
	right := 1 + solveDiameter(root.Right, edge, ans)
	*ans = util.Max(*ans, left+right)

	return util.Max(left, right)
}
