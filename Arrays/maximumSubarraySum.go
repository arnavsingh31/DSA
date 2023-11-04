package arrays

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #53
TC--->O(n^2)
SC--->O(1)
*/
func MaxSubarraySum(nums []int) int {
	max_sum := math.MinInt

	for i := 0; i < len(nums); i++ {
		curr_sum := 0
		for j := i; j < len(nums); j++ {
			curr_sum += nums[j]
			max_sum = util.Max(max_sum, curr_sum)
		}
	}

	return max_sum
}

/*
LC #53
TC-->O(n)
SC--->O(1)
Kadane's algorithm
*/
func MaxSubarraySum2(nums []int) int {
	curr_sum := 0
	max_sum := math.MinInt

	for i := 0; i < len(nums); i++ {
		curr_sum += nums[i]
		max_sum = util.Max(max_sum, curr_sum)

		if curr_sum < 0 {
			curr_sum = 0
		}
	}

	return max_sum
}
