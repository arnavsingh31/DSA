package binarysearch

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #1482
TC--->O(n + x*n) {TLE}
SC--->O(1)
Brute Force
*/
func MinDaysToMakeMBouquets(bloomDay []int, m, k int) int {
	n := len(bloomDay)
	// case where it is not possible to make m bouquets with k flowers
	if n/k < m {
		return -1
	}

	min := math.MaxInt
	max := math.MinInt

	// O(n) to find range of possible answers
	for i := 0; i < n; i++ {
		min = util.Min(min, bloomDay[i])
		max = util.Max(max, bloomDay[i])
	}

	//O(x) iterating over range x = (max - min)
	for i := min; i <= max; i++ {
		isPossible := possibleToMakeMBouquets(bloomDay, m, k, i)

		if isPossible {
			return i
		}
	}

	return -1
}

func possibleToMakeMBouquets(bloomDay []int, m, k, day int) bool {
	count := 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= day {
			// flower has bloomed
			count++
		} else {
			m -= count / k
			count = 0
		}
	}

	m -= count / k

	return m <= 0
}

/*
TC--->O( n + nlogn)~nlogn
SC--->O(1)
Using Binary search on possible answers
*/
func MinDaysToMakeMBouquets2(bloomDay []int, m, k int) int {
	n := len(bloomDay)
	// case where it is not possible to make m bouquets with k flowers
	if n/k < m {
		return -1
	}

	left := math.MaxInt
	right := math.MinInt

	// O(n) to find range of possible answers
	for i := 0; i < n; i++ {
		left = util.Min(left, bloomDay[i])
		right = util.Max(right, bloomDay[i])
	}

	//O(x) iterating over range x = (max - min) using binary search
	ans := right
	for left <= right {
		mid := (left + right) / 2

		if possibleToMakeMBouquets(bloomDay, m, k, mid) {
			ans = util.Min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
