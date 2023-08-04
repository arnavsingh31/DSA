package main

/*
	TC--->O(n)
	SC--->O(n)
	recursive
*/
func maxRootToLeafPathSum(root *Node) int {
	return helper(root, 0)
}

func helper(root *Node, sum int) int {
	if root == nil {
		return sum
	}

	leftPathSum := helper(root.Left, root.Val+sum)
	rightPathSum := helper(root.Right, root.Val+sum)
	return maxSum(leftPathSum, rightPathSum)
}

func maxSum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
