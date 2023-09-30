package trees

/*
	TC-->O(n)
	SC-->O(1)
*/
func InorderSuccPred(root *Node, data int) (int, int) {
	pred, succ := -1, -1
	temp := root
	for temp.Val != data {
		if temp.Val > data {
			succ = temp.Val
			temp = temp.Left
		} else {
			pred = temp.Val
			temp = temp.Right
		}
	}

	leftSubTree := temp.Left
	for leftSubTree != nil {
		pred = leftSubTree.Val
		leftSubTree = leftSubTree.Right
	}

	rightSubTree := temp.Right
	for rightSubTree != nil {
		succ = rightSubTree.Val
		rightSubTree = rightSubTree.Left
	}

	return pred, succ
}
