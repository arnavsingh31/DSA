package trees

/*
	GFG problem
*/
func BoundaryTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	ans = append(ans, root.Val)

	// traverse and store all left nodes
	traverseLeftNode(root.Left, &ans)

	// traverse all leaf nodes of left subtree
	traverseLeafNode(root.Left, &ans)

	// traverse all leaf nodes of right subtree
	traverseLeafNode(root.Right, &ans)

	//tarverse all right nodes
	traverseRightNode(root.Right, &ans)

	return ans
}

func traverseLeftNode(root *Node, ans *[]int) {
	if (root == nil) || (root.Left == nil && root.Right == nil) {
		return
	}

	*ans = append(*ans, root.Val)

	if root.Left != nil {
		traverseLeftNode(root.Left, ans)
	} else {
		traverseLeftNode(root.Right, ans)
	}
}

func traverseLeafNode(root *Node, ans *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, root.Val)
		return
	}

	traverseLeafNode(root.Left, ans)
	traverseLeafNode(root.Right, ans)
}

func traverseRightNode(root *Node, ans *[]int) {
	if (root == nil) || (root.Left == nil && root.Right == nil) {
		return
	}

	if root.Right != nil {
		traverseRightNode(root.Right, ans)
	} else {
		traverseRightNode(root.Left, ans)
	}

	*ans = append(*ans, root.Val)
}
