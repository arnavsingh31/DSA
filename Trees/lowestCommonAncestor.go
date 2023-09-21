package trees

/*
	LC #236
	TC---> O(n)
	SC---> O(n)
*/
func LowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
