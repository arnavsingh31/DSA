package binarysearch

import (
	"math"
)

/*
GFG problem
TC--->O(logn)
SC--->O(1)
if we can
*/
func TimesArrayRotated(arr []int) int {
	left, right := 0, len(arr)-1
	ans := math.MaxInt // gives the smallest element in the rotated sorted array.
	index := -1
	for left <= right {
		mid := (left + right) / 2

		/*
			In this case, the array from index low to high is completely sorted. Therefore, we can select the minimum element, arr[low].
			Now, if arr[low] < ans, we will update ‘ans’ with the value arr[low] and ‘index’ with the corresponding index low.
		*/
		if arr[left] <= arr[right] {
			if arr[left] < ans {
				ans = arr[left]
				index = left
				break
			}
		}

		if arr[mid] >= arr[left] {
			if arr[left] < ans {
				index = left
				ans = arr[left]
			}

			left = mid + 1
		} else {
			if arr[mid] < ans {
				index = mid
				ans = arr[mid]
			}

			right = mid - 1
		}
	}

	return index
}
