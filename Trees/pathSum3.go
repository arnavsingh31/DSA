package trees

/*
	LC #437
	TC--->O(n^2)
	SC--->O(logn)
*/
func PathSum3(root *Node, targetSum int) int {
	res := 0
	preOrder(root, targetSum, &res)
	return res
}

func preOrder(root *Node, targetSum int, res *int) {
	if root == nil {
		return
	}
	helperPathSum3(root, targetSum, res)
	preOrder(root.Left, targetSum, res)
	preOrder(root.Right, targetSum, res)
}

func helperPathSum3(root *Node, targetSum int, res *int) {
	if root == nil {
		return
	}

	if targetSum == root.Val {
		*res++
	}

	targetSum = targetSum - root.Val
	helperPathSum3(root.Left, targetSum, res)
	helperPathSum3(root.Right, targetSum, res)
}

/*
	TC--->O(n)
	SC--->O(n)
	optimised using hashmap to store sum from root to each node.
*/
func PathSum3_2(root *Node, targetSum int) int {
	pathSumMap := make(map[int]int)
	pathSumMap[0] = 1
	res := 0

	var helper func(*Node, int)
	helper = func(root *Node, currSum int) {
		if root == nil {
			return
		}
		currSum += root.Val
		if val, exist := pathSumMap[currSum-targetSum]; exist {
			res += val
		}

		pathSumMap[currSum]++
		helper(root.Left, currSum)
		helper(root.Right, currSum)
		pathSumMap[currSum]--
	}

	helper(root, 0)

	return res
}
