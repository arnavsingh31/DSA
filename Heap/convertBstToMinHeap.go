package heap

/*
	GFG
	BST is complete binary tree {Given}.
	TC--->O(n)
	SC--->O(n)
*/
func ConvertBST(root *Node) *Node {
	arr := inorder(root)

	// for min-heap, root < left-child && root < right-child; we want left-child < right-child
	// so combining both relation we get root < left-child < right-child.
	preorder(root, arr, 0)

	return root
}

// left -> root -> right
func inorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	left := inorder(root.Left)
	right := inorder(root.Right)

	return append(append(append([]int{}, left...), root.Val), right...)
}

func preorder(root *Node, arr []int, i int) {
	if root == nil {
		return
	}

	root.Val = arr[i]
	preorder(root.Left, arr, i+1)
	preorder(root.Right, arr, i+1)
}
