package main

/*
	LC #114
	TC---> O(n) traversing through each node twice.
	SC---> O(n) for call stack and storing nodes in array.
*/
func flatten(root *Node) {
	if root == nil {
		return
	}
	visited := preOrderTraversal2(root)
	prev := visited[0]
	visited = visited[1:]

	for _, node := range visited {
		prev.Left = nil
		prev.Right = node
		prev = node
	}
	prev.Left = nil
	prev.Right = nil
}

func preOrderTraversal2(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}
	leftNodes := preOrderTraversal2(root.Left)
	rightNodes := preOrderTraversal2(root.Right)

	return append(append(append([]*Node{}, root), leftNodes...), rightNodes...)
}

/*
	using df traversal with extreme right node first.
*/
func flatten2(root *Node) {
	var prevNode *Node
	var flatHelper func(*Node)

	flatHelper = func(node *Node) {
		if node != nil {
			flatHelper(node.Right)
			flatHelper(node.Left)
			node.Right = prevNode
			node.Left = nil
			prevNode = node
		}
	}
	flatHelper(root)
}

/*
	 1
    /  \
    2   5
   / \   \
  3   4   6

*/
