package binarysearch

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #875
TC--->O(nlogn)
SC--->O(1)
*/
func MinEatingSpeed(piles []int, h int) int {
	max := 0
	for i := 0; i < len(piles); i++ {
		max = util.Max(max, piles[i])
	}

	left, right := 1, max

	for left <= right {
		mid := (left + right) / 2

		timeReq := timeTaken(piles, mid)

		if timeReq <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right + 1
}

func timeTaken(arr []int, k int) int {
	time := 0
	for i := 0; i < len(arr); i++ {
		time += int(math.Ceil(float64(arr[i]) / float64(k)))
	}
	return time
}
