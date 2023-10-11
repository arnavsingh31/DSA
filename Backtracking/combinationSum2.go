package backtracking

import "sort"

/*
	LC #40
	TC--->O(2^n + nlogn{sorting})
*/
func CombinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	solveRec(candidates, target, 0, []int{}, &ans)
	return ans
}

func solveRec(arr []int, target, index int, curr []int, ans *[][]int) {
	if target == 0 {
		temp := append([]int{}, curr...)
		*ans = append(*ans, temp)
		return
	}

	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}
		if arr[i] > target {
			break
		}
		curr = append(curr, arr[i])
		solveRec(arr, target-arr[i], i+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
