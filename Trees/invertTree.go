package trees

/*
	LC #226
	TC---> O(n)
	SC---> O(n)
*/
func InvertTree(root *Node) *Node {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
