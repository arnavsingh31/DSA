package trees

/*
	LC #113
	TC--> O(n)
	SC--> O(n)
*/
func PathSum(root *Node, targetSum int) [][]int {
	res := [][]int{}
	findPath(root, targetSum, []int{}, &res)
	return res
}

func findPath(root *Node, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		*res = append(*res, append([]int{}, path...))
		return
	}

	findPath(root.Left, targetSum-root.Val, path, res)
	findPath(root.Right, targetSum-root.Val, path, res)
}
