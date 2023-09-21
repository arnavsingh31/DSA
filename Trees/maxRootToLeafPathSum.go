package trees

import util "github.com/arnavsingh31/DSA/Util"

/*
	TC--->O(n)
	SC--->O(n)
	recursive
*/
func MaxRootToLeafPathSum(root *Node) int {
	return helper(root, 0)
}

func helper(root *Node, sum int) int {
	if root == nil {
		return sum
	}

	leftPathSum := helper(root.Left, root.Val+sum)
	rightPathSum := helper(root.Right, root.Val+sum)
	return util.Max(leftPathSum, rightPathSum)
}
