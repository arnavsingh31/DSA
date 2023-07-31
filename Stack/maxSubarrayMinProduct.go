package main

import (
	"log"
	"math"
)

func maxSumMinProduct(nums []int) int {
	MOD := int(math.Pow10(9) + 7)
	inc_stack := make([]int, 0)
	prefixSum := []int{0}
	res := 0

	for i := 0; i < len(nums); i++ {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+nums[i])
	}

	for i := 0; i < len(nums); i++ {
		for len(inc_stack) > 0 && nums[inc_stack[len(inc_stack)-1]] > nums[i] {

			minVal := nums[inc_stack[len(inc_stack)-1]]
			inc_stack = inc_stack[:len(inc_stack)-1]

			prev_index := -1
			if len(inc_stack) > 0 {
				prev_index = inc_stack[len(inc_stack)-1]
			}

			rangeSum := prefixSum[i] - prefixSum[prev_index+1]
			res = maxSum(res, rangeSum*minVal)
		}
		inc_stack = append(inc_stack, i)
	}

	for i := 0; i < len(inc_stack); i++ {
		minVal := nums[inc_stack[i]]
		rangeSum := prefixSum[len(nums)] - prefixSum[inc_stack[i]]
		res = maxSum(res, rangeSum*minVal)
	}
	log.Print(res)
	return res % MOD
}

func maxSum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// maxSumMinProduct([]int{1, 2, 3, 2})
	// maxSumMinProduct([]int{3, 1, 5, 6, 4, 2})
	// maxSumMinProduct([]int{2, 3, 3, 1, 2})
	maxSumMinProduct([]int{4, 10, 6, 4, 8, 7, 8, 3, 5, 3, 4, 9, 9, 5, 10, 7, 10, 7, 6, 4})
}
