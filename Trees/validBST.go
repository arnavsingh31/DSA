package trees

import "math"

func IsValidBST(root *Node) bool {

	return checkBST(root, math.MinInt, math.MaxInt)

}

func checkBST(root *Node, minRange, maxRange int) bool {
	if root == nil {
		return true
	}

	if root.Val > minRange && root.Val < maxRange {
		left := checkBST(root.Left, minRange, root.Val)
		right := checkBST(root.Right, root.Val, maxRange)

		return left && right
	} else {
		return false
	}

}
