package trees

/*
	LC #100
	TC---> O(n)
	SC---> O(n)
*/
func IsSameTree(p, q *Node) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
