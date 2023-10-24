package main

/*
	LC #1673
	TC--->O(n)
	SC--->O(n) (if we exclude result array which is same stack then complexity reduces to O(1))
	Using monotonic increasing stack
*/
func mostCompetitveSubsequence(arr []int, k int) []int {
	inc_stack := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		for len(inc_stack) > 0 && (inc_stack[len(inc_stack)-1] > arr[i]) && k-len(inc_stack) < len(arr)-i {
			inc_stack = inc_stack[:len(inc_stack)-1]
		}
		if len(inc_stack) < k {
			inc_stack = append(inc_stack, arr[i])
		}
	}

	return inc_stack
}
