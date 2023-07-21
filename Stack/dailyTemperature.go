package main

/*
	LC #739
	Using monotonic stack we have to find nextgreater temperature index and store difference of current index
	and next higher index in waitDay array.
	T.C--> O(n), S.C---> O(n)
*/
func dailyTemperature(temp []int) []int {
	waitDay := []int{}
	stack := []int{}

	for i := 0; i < len(temp); i++ {
		waitDay = append(waitDay, 0)
	}

	for i, t := range temp {
		for len(stack) > 0 && temp[stack[len(stack)-1]] < t {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			waitDay[stackTop] = i - stackTop
		}
		stack = append(stack, i)
	}

	return waitDay

}