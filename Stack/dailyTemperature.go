package stack

/*
	LC #739
	Using monotonic stack we have to find nextgreater temperature index and store difference of current index
	and next higher index in waitDay array.
	T.C--> O(n), S.C---> O(n)
*/
func DailyTemperature(temp []int) []int {
	waitDay := make([]int, len(temp))
	stack := []int{}

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
