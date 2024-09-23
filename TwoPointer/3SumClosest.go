package twopointer

import (
	"math"
	"sort"
)

// LC #16
// TC := O(N^2)
// SC := O(1)
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt
	n := len(nums) - 1
	var closestSum int

	for i := 0; i < n; i++ {
		if i == n-1 {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == target {
				return sum
			}

			tempDiff := int(math.Abs(float64(target - sum)))

			if tempDiff < diff {
				diff = tempDiff
				closestSum = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return closestSum
}
