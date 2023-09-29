package trees

/*
	LC #450
	TC--->O(n)
	SC--->O(n)
*/
func DeleteNodeInBST(root *Node, data int) *Node {
	// if node we want to delete does not exist in tree
	if root == nil {
		return nil
	}

	// found the node
	if root.Val == data {
		// node is leaf node
		if root.Left == nil && root.Right == nil {
			root = nil
			return root
		}

		// node has single child/subtree
		if root.Left != nil && root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}

		if root.Left == nil && root.Right != nil {
			temp := root.Right
			root = nil
			return temp
		}

		// node has 2 child/ subtree
		if root.Left != nil && root.Right != nil {
			// get max value from left subtree, copy its value to node and then delete that max node from leftsubtree.
			max := getMaxNodeVal(root.Left).Val
			root.Val = max
			root.Left = DeleteNodeInBST(root.Left, max)
			return root

			// alternative get min value from right subtree, copy its value to node and then delete that min node from right subtree.
			// min := getMinNodeVal(root.Right).Val
			// root.Val = min
			// root.Right = DeleteNodeInBST(root.Right, min)
			// return root
		}

	} else if root.Val > data {
		root.Left = DeleteNodeInBST(root.Left, data)
	} else {
		root.Right = DeleteNodeInBST(root.Right, data)
	}

	return root
}

func getMaxNodeVal(root *Node) *Node {
	if root.Right == nil {
		return root
	}
	return getMaxNodeVal(root.Right)
}

func getMinNodeVal(root *Node) *Node {
	if root.Left == nil {
		return root
	}
	return getMinNodeVal(root.Left)
}
