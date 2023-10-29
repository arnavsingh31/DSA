package slidingwindow

import util "github.com/arnavsingh31/DSA/Util"

// my solution sliding window.
func MinSubArrayLen(target int, arr []int) int {
	if len(arr) == 1 {
		if arr[0] >= target {
			return 1
		} else {
			return 0
		}
	}

	left, right := 0, 0
	minLen := len(arr) + 1
	sum := arr[left] // initial sum

	for left <= right {
		if sum < target && right+1 <= len(arr)-1 {
			sum += arr[right+1]
			right++
		} else if sum >= target {

			minLen = util.Min(right-left+1, minLen)
			sum -= arr[left]
			left++
		} else {
			break
		}
	}

	if minLen == len(arr)+1 {
		return 0
	}

	return minLen
}

// actual way to implement sliding window algorithm. Note: ***T.C is O(n) not O(n^2)
// Read from this for recalling ---> https://leetcode.com/explore/featured/card/leetcodes-interview-crash-course-data-structures-and-algorithms/703/arraystrings/4502/
func MinSubArrayLen2(target int, arr []int) int {
	left, sum, minLen := 0, 0, len(arr)+1

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		for sum >= target {
			minLen = util.Min(right-left+1, minLen)
			sum -= arr[left]
			left++
		}
	}

	if minLen == len(arr)+1 {
		return 0
	}
	return minLen
}
