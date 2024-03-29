package stack

func PreviousGreater(arr []int) []int {
	stack := []int{}
	prevGr := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		prevGr = append(prevGr, -1)
	}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			prevGr[i] = arr[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return prevGr
}
