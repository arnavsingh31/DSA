package trees

/*
	LC #103
	TC--> O(n)
	SC--> O(n)
*/
func ZigZagTraversal(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*Node{root}
	ans := [][]int{}
	reverse := false

	for len(queue) > 0 {
		levelLength := len(queue)
		arr := make([]int, 0)

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if reverse {
				arr = append(append([]int{}, node.Val), arr...)
			} else {
				arr = append(arr, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		reverse = !reverse
		ans = append(ans, arr)
	}
	return ans
}
