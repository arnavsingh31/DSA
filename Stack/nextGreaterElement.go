package main

// LC #496
func nextGreaterElement(nums1, nums2 []int) []int {
	nextGr := map[int]int{}
	stack := []int{}

	for _, num := range nums2 {
		nextGr[num] = -1
		for len(stack) > 0 && stack[len(stack)-1] < num {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGr[stackTop] = num
		}
		stack = append(stack, num)
	}

	res := []int{}
	for _, num := range nums1 {
		res = append(res, nextGr[num])
	}

	return res
}
