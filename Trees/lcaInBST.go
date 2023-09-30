package trees

/*
	TC--->O(n)
	SC--->O(n) {call stack}--> can be reduced to O(1) by iterative approach.
*/
func LCAInBST(root *Node, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Val < p.Val && root.Val < q.Val {
		LCAInBST(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		LCAInBST(root.Left, p, q)
	}

	return root
}
