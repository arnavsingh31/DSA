package trees

/*
	LC #112
	TC---> O(n)
	SC--->O(n)
*/
func HasPathSum(root *Node, target int) bool {
	if root == nil {
		return false
	}

	target -= root.Val

	if root.Left == nil && root.Right == nil {
		return target == 0
	}

	return HasPathSum(root.Left, target) || HasPathSum(root.Right, target)
}
