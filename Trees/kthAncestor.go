package trees

import "math"

/*
	GFG problem

*/
func FindKthAncestor(root *Node, node, k int) *Node {
	ans := recHelper(root, node, &k)

	if root == nil || ans.Val == node {
		return nil
	}
	return ans
}

func recHelper(root *Node, node int, k *int) *Node {
	if root == nil {
		return nil
	}

	if root.Val == node {
		// *k--
		return root
	}

	left := recHelper(root.Left, node, k)
	right := recHelper(root.Right, node, k)

	// if *k == 0 {
	// 	*k--
	// 	return root
	// }

	// if left != nil || right != nil {
	// 	*k--
	// }

	// if left != nil {
	// 	return left
	// }

	// return right

	if left != nil && right == nil {
		*k--
		if *k <= 0 {
			*k = math.MaxInt
			return root
		}
		return left
	}

	if left == nil && right != nil {
		*k--
		if *k <= 0 {
			*k = math.MaxInt
			return root
		}
		return right
	}

	return nil
}
