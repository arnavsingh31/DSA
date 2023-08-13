package main

/*
	LC #236
	TC---> O(n)
	SC---> O(n)
*/
func lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
