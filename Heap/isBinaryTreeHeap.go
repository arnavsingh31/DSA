package heap

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*
	TC--->O(n)
	SC--->O(n)
	Here for a Binary tree to be a maxheap must follow 2 conditons:
	1. It must be a CBT
	2. It must satify conditons of a max-heap.
*/

func IsHeap(root *Node) bool {

	if isCompleteBinaryTree2(root, 0, countNodes(root)) && isMaxHeap(root) {
		return true
	}

	return false
}

// Method-1 using level traversal
func isCompleteBinaryTree(node *Node) bool {
	queue := []*Node{node}
	flag := false

	for len(queue) > 0 {
		levelLength := len(queue)

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				if flag {
					return false
				}
				queue = append(queue, node.Left)
			} else {
				flag = true
			}

			if node.Right != nil {
				if flag {
					return false
				}
				queue = append(queue, node.Right)
			} else {
				flag = true
			}
		}
	}

	return true
}

// Method-2 using index property
func isCompleteBinaryTree2(root *Node, index int, totalNodes int) bool {
	if root == nil {
		return true
	}

	if index >= totalNodes {
		return false
	}
	left := isCompleteBinaryTree2(root.Left, 2*index+1, totalNodes)
	right := isCompleteBinaryTree2(root.Right, 2*index+2, totalNodes)
	return left && right
}

func countNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

func isMaxHeap(node *Node) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return true
	}

	if node.Right == nil {
		return node.Val >= node.Left.Val
	} else {
		if node.Val >= node.Left.Val && node.Val >= node.Right.Val {
			left := isMaxHeap(node.Left)
			right := isMaxHeap(node.Right)

			return left && right
		} else {
			return false
		}
	}
}
