package trees

/*
	LC #101
	TC---> O(n)
	SC---> O(n)
*/
func IsSymmetric(root *Node) bool {
	return symHelper(root.Left, root.Right)
}

func symHelper(r1, r2 *Node) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil || r1.Val != r2.Val {
		return false
	}

	return symHelper(r1.Left, r2.Right) && symHelper(r1.Right, r2.Left)

}
