package trees

/*
	Morris traversal runs in constant space complexity, and linear time complexity.
	inorder tarversal
*/
func MorrisInorderTraversal(root *Node) []int {
	ans := make([]int, 0)

	curr := root

	for curr != nil {
		if curr.Left == nil {
			ans = append(ans, curr.Val)
			curr = curr.Right
		} else {
			prev := curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
			} else {
				ans = append(ans, curr.Val)
				prev.Right = nil
				curr = curr.Right
			}
		}
	}

	return ans
}

// preorder traversal
func MorrisPreorderTraversal(root *Node) []int {
	ans := make([]int, 0)

	curr := root

	for curr != nil {
		if curr.Left == nil {
			ans = append(ans, curr.Val)
			curr = curr.Right
		} else {
			prev := curr.Left

			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = curr
				ans = append(ans, curr.Val)
				curr = curr.Left
			} else {
				prev.Right = nil
				curr = curr.Right
			}
		}
	}
	return ans
}
