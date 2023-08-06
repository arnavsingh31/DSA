package main

/*
	LC #226
	TC---> O(n)
	SC---> O(n)
*/
func invertTree(root *Node) *Node {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
