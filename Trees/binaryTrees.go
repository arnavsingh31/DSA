package trees

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
TC--> O(n)
SC--> O(n)
Depth First Traversal(DFS) Algorithm using stack
*/
func TreeDepthFirstTraversal(root *Node) []int {
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
func TreeDepthFirstTraversal2(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftNodes := TreeDepthFirstTraversal(root.Left)
	rightNodes := TreeDepthFirstTraversal(root.Right)

	ans := []int{root.Val}
	ans = append(ans, leftNodes...)
	ans = append(ans, rightNodes...)
	return ans
}

/*
TC--> O(n)
SC--> O(n)
*/
func BreadFirstTraversal(root *Node) []int {
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
func Bfs(root *Node, val int) bool {
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
func Dfs(root *Node, val int) bool {
	// base cases
	if root == nil {
		return false
	}

	if root.Val == val {
		return true
	}

	return Dfs(root.Left, val) || Dfs(root.Right, val)
}

// func main() {
// root := &Node{Val: 1}
// node1 := &Node{Val: 2}
// node2 := &Node{Val: 3}
// node3 := &Node{Val: 4}
// node4 := &Node{Val: 5}
// node5 := &Node{Val: 6}

// root.Left = node1
// root.Right = node2
// node1.Left = node3
// node1.Right = node4
// node2.Right = node5

// log.Print(breadFirstTraversal(root))
// log.Print(searchValue2(root, 4))
// log.Print(treeSum2(root))
// log.Print(treeMin(root))
// log.Print(inorderTraversal(root))
// log.Print(preOrderTraversal(root))
// log.Print(postOrderTraversal(root))
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

// Iterative approach
func TreeSum(root *Node) int {
	stack := []*Node{root}
	sum := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += node.Val

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return sum
}

// recursive approach
func TreeSum2(root *Node) int {
	if root == nil {
		return 0
	}

	sum := root.Val + TreeSum2(root.Left) + TreeSum2(root.Right)
	return sum
}

/*
Finding min value in a binary tree.
Iterative Approach
*/
func TreeMin(root *Node) int {
	stack := []*Node{root}
	min := math.MaxInt

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val < min {
			min = node.Val
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return min
}

func TreeMin2(root *Node) int {
	if root == nil {
		return math.MaxInt
	}

	min := util.Min(root.Val, TreeMin2(root.Left))
	min = util.Min(min, TreeMin2(root.Right))

	return min
}

/*
LC #94
left -> root -> right
*/
func InorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftNodes := InorderTraversal(root.Left)
	rightNodes := InorderTraversal(root.Right)

	return append(append(append([]int{}, leftNodes...), root.Val), rightNodes...)
}

// root -> left -> right (same as depth first traversal)
func PreOrderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftNodes := PreOrderTraversal(root.Left)
	rightNodes := PreOrderTraversal(root.Right)

	return append(append(append([]int{}, root.Val), leftNodes...), rightNodes...)
}

// left -> right -> root
func PostOrderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	leftNodes := PostOrderTraversal(root.Left)
	rightNodes := PostOrderTraversal(root.Right)

	return append(append(append([]int{}, leftNodes...), rightNodes...), root.Val)
}
