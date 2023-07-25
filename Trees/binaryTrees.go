package main

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

/*
TC--> O(n)
SC--> O(n)
Depth First Search(DFS) Algorithm using stack
*/
func dfs(root *Node) []int {
	stack := []*Node{root}
	visited := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited = append(visited, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return visited
}

/*
DFS using recursion
*/
func dfs2(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftNodes := dfs2(root.Left)
	rightNodes := dfs2(root.Right)

	ans := []int{root.Val}
	ans = append(ans, leftNodes...)
	ans = append(ans, rightNodes...)
	return ans
}

// func main() {
// 	root := &Node{Val: 1}
// 	node1 := &Node{Val: 2}
// 	node2 := &Node{Val: 3}
// 	node3 := &Node{Val: 4}
// 	node4 := &Node{Val: 5}
// 	node5 := &Node{Val: 6}

// 	root.Left = node1
// 	root.Right = node2
// 	node1.Left = node3
// 	node1.Right = node4
// 	node2.Right = node5

// 	log.Print(dfs2(root))
// }

/*
	 1
    /  \
    2   3
   / \   \
  4   5   6
*/
