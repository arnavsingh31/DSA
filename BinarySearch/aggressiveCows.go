package binarysearch

import (
	"math"
	"sort"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
GFG Aggressive Cows
TC--->O(nlogn{sorting} + nlog(maxPosStall - minPosStall))
SC--->O(1)
Using Binary Search on answers
*/
func MaxOfMinDistanceBetweenCows(stalls []int, k int) int {
	sort.Ints(stalls)
	left, right := stalls[0], stalls[len(stalls)-1]-stalls[0]
	ans := math.MinInt

	for left <= right {
		mid := (left + right) / 2

		if distanceBetweenCowsPossible(stalls, mid, k) {
			ans = util.Max(ans, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func distanceBetweenCowsPossible(stalls []int, dis, k int) bool {
	prevCowPos := stalls[0]

	for i := 1; i < len(stalls); i++ {
		currPos := stalls[i]

		if currPos-prevCowPos >= dis {
			k--
			prevCowPos = currPos
		}
	}

	return k <= 0
}
