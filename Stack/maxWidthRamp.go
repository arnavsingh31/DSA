package main

import "log"

/*
	LC #962
	TC---> O(n^2)
	SC---> O(1)
	Bruteforce approach
*/
func maxWidthRamp(arr []int) int {
	max := -1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] {
				max = maxRamp(max, j-i)
			}
		}
	}
	return max
}

/*
	TC---> O(n)
	SC---> O(n)
	Monotonic stack approach
*/

func maxRamp(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWidthRamp2(arr []int) int {
	max := -1
	stack := make([]int, 0)

	for i, num := range arr {
		if i == 0 || arr[stack[len(stack)-1]] > num {
			stack = append(stack, i)
		}
	}

	for j := len(arr) - 1; j >= 0; j-- {
		for len(stack) > 0 && arr[j] >= arr[stack[len(stack)-1]] {
			max = maxRamp(max, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return max
}

func main() {
	log.Print(maxWidthRamp2([]int{6, 0, 8, 2, 1, 5}))
}
