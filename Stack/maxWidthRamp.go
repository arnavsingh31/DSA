package stack

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #962
	TC---> O(n^2)
	SC---> O(1)
	Bruteforce approach
*/
func MaxWidthRamp(arr []int) int {
	max := -1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] {
				max = util.Max(max, j-i)
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

func MaxWidthRamp2(arr []int) int {
	max := -1
	stack := make([]int, 0)

	for i, num := range arr {
		if i == 0 || arr[stack[len(stack)-1]] > num {
			stack = append(stack, i)
		}
	}

	for j := len(arr) - 1; j >= 0; j-- {
		for len(stack) > 0 && arr[j] >= arr[stack[len(stack)-1]] {
			max = util.Max(max, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return max
}

// func main() {
// 	log.Print(maxWidthRamp2([]int{6, 0, 8, 2, 1, 5}))
// }
