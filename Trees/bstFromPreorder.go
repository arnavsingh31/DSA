package trees

import "math"

func BSTFromPreorder(arr []int) *Node {
	index := 0
	root := makeBST(arr, math.MaxInt, math.MinInt, &index)

	return root
}

func makeBST(arr []int, upper, lower int, index *int) *Node {
	if *index >= len(arr) {
		return nil
	}

	if arr[*index] < lower || arr[*index] > upper {
		return nil
	}

	node := &Node{Val: arr[*index]}
	*index += 1
	node.Left = makeBST(arr, node.Val, lower, index)
	node.Right = makeBST(arr, upper, node.Val, index)
	return node
}
