package main

/*
	LC #1475
	TC--->O(n)
	SC--->O(n)
	Using monotonic increasing stack
*/
func finalPrice(prices []int) []int {
	ans := make([]int, len(prices))
	inc_stack := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for len(inc_stack) > 0 && prices[inc_stack[len(inc_stack)-1]] >= prices[i] {
			pop_index := inc_stack[len(inc_stack)-1]
			inc_stack = inc_stack[:len(inc_stack)-1]

			ans[pop_index] = prices[pop_index] - prices[i]
		}
		inc_stack = append(inc_stack, i)
	}

	for _, pos := range inc_stack {
		ans[pos] = prices[pos]
	}

	return ans
}
