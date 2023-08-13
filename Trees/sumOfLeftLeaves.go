package main

/*
	LC #404
	TC--->O(n)
	SC--->O(n)
*/
func sumOfLeftLeaves(root *Node) int {
	return helperLeftLeaf(root, false)
}

func helperLeftLeaf(root *Node, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		}
		return 0
	}

	leftVal := helperLeftLeaf(root.Left, true)
	rightVal := helperLeftLeaf(root.Right, false)

	return leftVal + rightVal
}
