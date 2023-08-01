package main

import "math"

/*
	LC #907
	TC--->O(n)
	SC--->O(n)
	Using monotonic increasing stack
*/
func sumSubarrayMins(arr []int) int {
	MOD := int(math.Pow10(9) + 7)
	sum := 0
	inc_stack := make([]int, len(arr))

	for i := 0; i <= len(arr); i++ {
		for len(inc_stack) > 0 && ((i == len(arr)) || arr[inc_stack[len(inc_stack)-1]] > arr[i]) {
			pop_index := inc_stack[len(inc_stack)-1]
			minVal := arr[pop_index]
			inc_stack = inc_stack[:len(inc_stack)-1]
			prev_index := -1
			if len(inc_stack) > 0 {
				prev_index = inc_stack[len(inc_stack)-1]
			}
			count := (pop_index - prev_index) * (i - pop_index)
			sum += count * minVal
		}
		inc_stack = append(inc_stack, i)
	}

	return sum % MOD
}
