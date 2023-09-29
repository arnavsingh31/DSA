package trees

import util "github.com/arnavsingh31/DSA/Util"

/*
	GFG problem
*/
type Pair struct {
	first  int // max sum by including nodes at current level
	second int // max sum by excluding nodes at current level
}

func MaxSum(root *Node) int {
	ans := solveRec(root)
	return util.Max(ans.first, ans.second)
}

func solveRec(root *Node) Pair {
	if root == nil {
		return Pair{0, 0}
	}

	left := solveRec(root.Left)
	right := solveRec(root.Right)

	var ans Pair
	ans.first = root.Val + left.second + right.second
	ans.second = util.Max(left.first, left.second) + util.Max(right.first, right.second)

	return ans
}
