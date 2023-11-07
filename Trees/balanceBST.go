package trees

/*
	LC #1382
	TC--->O(n)
	SC--->O(n)
	combination of two questions,
	1. make inorder traversal of BST which will give sorted array
	2. usking above sorted array make a balance BST
*/
func BalanceBST(root *Node) *Node {
	arr := InorderTraversal(root)
	return buildBalancedBST(arr, 0, len(arr)-1)
}
