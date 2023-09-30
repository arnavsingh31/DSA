package trees

/*
	TC--->O(n)
	SC--->O(n)
*/
func NormalToBalancedBST(root *Node) *Node {
	arr := InorderTraversal(root)
	return buildBalancedBST(arr, 0, len(arr)-1)
}

func buildBalancedBST(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &Node{Val: arr[mid]}
	root.Left = buildBalancedBST(arr, start, mid-1)
	root.Right = buildBalancedBST(arr, mid+1, end)

	return root
}
