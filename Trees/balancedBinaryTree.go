package trees

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #110
TC--->O(n)
SC--->O(n)
*/
func IsBalanced(root *TreeNode) bool {
	return recur(root).isBal
}

type Pair2 struct {
	height int
	isBal  bool
}

func recur(root *TreeNode) Pair2 {
	if root == nil {
		return Pair2{0, true}
	}

	if root.Left == nil && root.Right == nil {
		return Pair2{1, true}
	}

	left := recur(root.Left)
	right := recur(root.Right)

	return Pair2{util.Max(left.height, right.height) + 1, int(math.Abs(float64(left.height-right.height))) <= 1 && left.isBal && right.isBal}
}
