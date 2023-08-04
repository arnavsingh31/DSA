package main

/*
	LC #129
	TC--->O(n)
	SC--->O(n)
*/
func sumNumbers(root *Node) int {
	return helperSum(root, 0)
}

func helperSum(root *Node, sum int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}

	leftPathSum := helperSum(root.Left, sum*10+root.Val)
	rightPathSum := helperSum(root.Right, sum*10+root.Val)

	return leftPathSum + rightPathSum
}
