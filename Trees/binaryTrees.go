package main

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

/*
TC--> O(n)
SC--> O(n)
Depth First Traversal(DFS) Algorithm using stack
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
Depth First Traversal using recursion
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

/*
TC--> O(n)
SC--> O(n)
*/
func breadFirstTraversal(root *Node) []int {
	queue := []*Node{root}
	ans := []int{}

	for len(queue) > 0 {
		node := queue[0]
		ans = append(ans, node.Val)
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return ans
}

// breadth first searching
func searchValue(root *Node, val int) bool {
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if node.Val == val {
			return true
		}
	}
	return false
}

// depth first searching (recursive)
func searchValue2(root *Node, val int) bool {
	// base cases
	if root == nil {
		return false
	}

	if root.Val == val {
		return true
	}

	return searchValue2(root.Left, val) || searchValue2(root.Right, val)
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

// 	log.Print(breadFirstTraversal(root))
// 	log.Print(searchValue2(root, 4))
// }

/*
	 1
    /  \
    2   3
   / \   \
  4   5   6

  Depth first traversal :- [1,2,4,5,3,6]
  Breadth First traversal :- [1,2,3,4,5,6]
*/
