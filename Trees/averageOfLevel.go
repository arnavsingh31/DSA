package main

/*
	LC #637
	TC---> O(n)
	SC---> O(n)
*/
func findAverage(root *Node) []float64 {
	queue := []*Node{root}
	ans := []float64{}
	for len(queue) > 0 {
		levelLength := len(queue)
		sum := 0.0
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, sum/float64(levelLength))
	}
	return ans
}
