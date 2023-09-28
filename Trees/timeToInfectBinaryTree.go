package trees

/*
	LC #2385
	TC--->O(n) --{O(n)-> to find startNode and populate child-parent mapping + O(n)--> bfs}
	SC--->O(n) -- {queue + child to parent map + infect node map}
*/
func AmountOfTime(root *Node, start int) int {
	timeTaken := 0
	queue := []*Node{root}
	childToParentMapping := make(map[*Node]*Node)
	infectNodes := make(map[*Node]bool)
	var startNode *Node

	// step 1 find starting node and also create child to parent mapping.
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val == start {
			startNode = node
		}

		if node.Left != nil {
			childToParentMapping[node.Left] = node
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			childToParentMapping[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	// step 2 do bfs from starting/infected node and calculate time
	queue = append(queue, startNode)
	infectNodes[startNode] = true
	for len(queue) > 0 {
		levelSize := len(queue)
		flag := false

		for i := 0; i < levelSize; i++ {

			node := queue[0]
			queue = queue[1:]

			if node.Left != nil && !infectNodes[node.Left] {
				queue = append(queue, node.Left)
				infectNodes[node.Left] = true
				flag = true
			}

			if node.Right != nil && !infectNodes[node.Right] {
				queue = append(queue, node.Right)
				infectNodes[node.Right] = true
				flag = true
			}

			if parent, exist := childToParentMapping[node]; exist && !infectNodes[parent] {
				queue = append(queue, parent)
				infectNodes[parent] = true
				flag = true
			}
		}
		if flag {
			timeTaken++
		}
	}

	return timeTaken
}
