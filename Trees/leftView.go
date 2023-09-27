package trees

func LeftView(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == 0 {
				ans = append(ans, node.Val)
			}

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

// recursive approach
func LeftView2(root *Node) []int {
	ans := make([]int, 0)
	rec(root, &ans, 0)
	return ans
}

func rec(root *Node, ans *[]int, level int) {
	if root == nil {
		return
	}

	if len(*ans) == level {
		*ans = append(*ans, root.Val)
	}

	rec(root.Left, ans, level+1)
	rec(root.Right, ans, level+1)
}
