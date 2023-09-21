package trees

/*
	LC #102
	TC--->O(n)
	SC--->O(n+ 2d array{ans})
*/
func LevelOrderTraversal(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*Node{root}
	ans := [][]int{}
	for len(queue) > 0 {
		levelLength := len(queue)
		levelNodes := []int{}

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		ans = append(ans, levelNodes)

	}
	return ans
}
