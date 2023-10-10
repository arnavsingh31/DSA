package backtracking

/*
	LC #39
	TC--->O(exponential)
	SC--->O()
*/
func CombinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	solveRecCom(candidates, target, []int{}, 0, &ans)
	return ans
}

func solveRecCom(arr []int, target int, curr []int, index int, ans *[][]int) {
	if target == 0 {
		temp := append([]int{}, curr...)
		*ans = append(*ans, temp)
		return
	}

	if index == len(arr) {
		return
	}

	if target >= arr[index] {
		// take current element
		curr = append(curr, arr[index])
		solveRecCom(arr, target-arr[index], curr, index, ans)

		// backtrack
		curr = curr[:len(curr)-1]
	}

	// not-take current element and move onto next index
	solveRecCom(arr, target, curr, index+1, ans)
}
