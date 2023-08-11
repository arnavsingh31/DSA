package main

/*
	LC #107
*/
func levelOrderTraversal2(root *Node) [][]int {
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
	// till here same as part1 of the problem now we just reverse the array

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}
