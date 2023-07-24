package main

/*
	LC #456
	TC---> O(n)
	SC--->O(n)
	Using monotonic stack. we use the logic of previous greater element to find index j and then use this j we
	find min[j] (which is our possible arr[i]) & if this min[j] < num (which is arr[k]) then we have a 132
	pattern and we return true.
*/
func is132Pattern(arr []int) bool {
	min := make([]int, len(arr))
	stack := []int{}

	for index, num := range arr {
		// calculating min value encountered till now for every index.
		if index == 0 {
			min[index] = num
		} else if num < min[index-1] {
			min[index] = num
		} else {
			min[index] = min[index-1]
		}

		for len(stack) > 0 && arr[stack[len(stack)-1]] <= num {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			if min[stack[len(stack)-1]] < num {
				return true
			}
		}

		stack = append(stack, index)
	}

	return false
}
