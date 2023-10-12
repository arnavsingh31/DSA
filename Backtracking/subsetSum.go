package backtracking

import "sort"

/*
	GFG-Problem
	TC--->O(2^n)
	SC--->O(2^n)
*/
func SubsetSum(arr []int) []int {
	ans := make([]int, 0)
	solveRecSubset(arr, 0, 0, &ans)
	sort.Ints(ans)
	return ans
}

func solveRecSubset(arr []int, sum, index int, ans *[]int) {
	if index >= len(arr) {
		*ans = append(*ans, sum)
		return
	}

	// take
	solveRecSubset(arr, sum+arr[index], index+1, ans)

	//not take
	solveRecSubset(arr, sum, index+1, ans)
}
