package main

/*
	LC #112
	TC---> O(n)
	SC--->O(n)
*/
func hasPathSum(root *Node, target int) bool {
	if root == nil {
		return false
	}

	target -= root.Val

	if root.Left == nil && root.Right == nil {
		return target == 0
	}

	return hasPathSum(root.Left, target) || hasPathSum(root.Right, target)
}
