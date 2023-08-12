package main

/*
	LC #199
	TC---> O(n)
	SC---> O(n)
*/
func rightSideView(root *Node) []int {
	if root == nil {
		return []int{}
	}

	queue := []*Node{root}
	ans := []int{}

	for len(queue) > 0 {
		levelLength := len(queue)

		for i := 0; i < levelLength; i++ {
			node := queue[0]

			if i == levelLength-1 {
				ans = append(ans, node.Val)
			}

			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}
