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
