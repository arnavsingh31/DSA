package binarysearch

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #1011
	TC--->O(n + (totalWeight-maxWeight)*n) {TLE}
	SC--->O(1)
*/
func CapacityToShipPackage(weights []int, days int) int {
	high := 0
	low := 0

	for i := 0; i < len(weights); i++ {
		high += weights[i]
		low = util.Max(low, weights[i])
	}

	for i := low; i <= high; i++ {
		daysReq := daysRequired(weights, i)

		if daysReq <= days {
			return i
		}
	}

	return -1
}

func daysRequired(weights []int, cap int) int {
	currLoad := 0
	daysReq := 1
	for i := 0; i < len(weights); i++ {
		if currLoad+weights[i] <= cap {
			currLoad += weights[i]
		} else {
			daysReq++
			currLoad = weights[i]
		}
	}

	return daysReq
}

/*
	TC--->O(n + log(totalWeight - maxWeight) * n)
	SC--->O(1)
	Using Binary Search on answers
*/
func CapacityToShipPackage2(weights []int, days int) int {
	high := 0
	low := 0
	for i := 0; i < len(weights); i++ {
		high += weights[i]
		low = util.Max(low, weights[i])
	}
	ans := high

	for low <= high {
		mid := (low + high) / 2
		daysReq := daysRequired(weights, mid)
		if daysReq <= days {
			ans = util.Min(ans, mid)
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}
