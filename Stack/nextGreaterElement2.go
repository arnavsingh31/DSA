package stack

// LC #503
func NextGreaterCircular(nums []int) []int {
	nextGr := []int{}
	stack := []int{}

	for i := 0; i < len(nums); i++ {
		nextGr = append(nextGr, -1)
	}

	for i := 0; i < 2; i++ {
		for index, num := range nums {
			for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
				stackTop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				nextGr[stackTop] = num
			}
			stack = append(stack, index)
		}

	}
	return nextGr
}
