package binarysearch

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #153
TC--->O(logn)
SC--->O(1)
*/
func MinInRotatedSortedArray(arr []int) int {
	left, right := 0, len(arr)-1
	ans := math.MaxInt

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] >= arr[left] {
			ans = util.Min(ans, arr[left])
			left = mid + 1
		} else {
			ans = util.Min(ans, arr[mid])
			right = mid - 1
		}
	}

	return ans
}
