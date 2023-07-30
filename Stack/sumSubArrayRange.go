package main

/*
	LC #2104
	TC---> O(n)
	SC ---> O(n)
*/
func subArrayRanges(nums []int) int64 {
	var sum int64
	sum = 0
	stack := make([]int, 0)

	for curr_index := 0; curr_index <= len(nums); curr_index++ {
		for len(stack) > 0 && ((curr_index == len(nums)) || nums[stack[len(stack)-1]] <= nums[curr_index]) {
			pop_Index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev_Index := -1
			if len(stack) > 0 {
				prev_Index = stack[len(stack)-1]
			}
			sum += int64((pop_Index - prev_Index) * (curr_index - pop_Index) * nums[pop_Index])
		}
		stack = append(stack, curr_index)
	}

	stack = []int{}

	for curr_index := 0; curr_index <= len(nums); curr_index++ {
		for len(stack) > 0 && ((curr_index == len(nums)) || nums[stack[len(stack)-1]] >= nums[curr_index]) {
			pop_Index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev_Index := -1
			if len(stack) > 0 {
				prev_Index = stack[len(stack)-1]
			}
			sum -= int64((pop_Index - prev_Index) * (curr_index - pop_Index) * nums[pop_Index])
		}
		stack = append(stack, curr_index)
	}

	return sum
}

// refactored code.
func subArrayRanges2(nums []int) int64 {
	var sum int64
	sum = 0
	dec_stack := make([]int, 0)
	inc_stack := make([]int, 0)

	for i := 0; i <= len(nums); i++ {
		sum += computeSum(&dec_stack, nums, i, func(i, j int) bool { return i <= j }) - computeSum(&inc_stack, nums, i, func(i, j int) bool { return i >= j })
	}

	return sum
}

func computeSum(stackPtr *[]int, nums []int, curr_index int, compare func(int, int) bool) int64 {
	var maxSum int64
	maxSum = 0
	stack := *stackPtr

	for len(stack) > 0 && ((curr_index == len(nums)) || compare(nums[stack[len(stack)-1]], nums[curr_index])) {
		pop_Index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev_Index := -1
		if len(stack) > 0 {
			prev_Index = stack[len(stack)-1]
		}
		maxSum += int64((pop_Index - prev_Index) * (curr_index - pop_Index) * nums[pop_Index])
	}

	stack = append(stack, curr_index)
	*stackPtr = stack

	return maxSum
}
