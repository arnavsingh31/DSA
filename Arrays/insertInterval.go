package arrays

import util "github.com/arnavsingh31/DSA/Util"

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	ans := make([][]int, 0)
	pos := 0

	for i := 0; i < n; i++ {
		currInterval := intervals[i]

		if currInterval[1] >= newInterval[0] {
			break
		}

		ans = append(ans, intervals[i])
		pos = i + 1
	}

	for i := pos; i < n; i++ {
		currInterval := intervals[i]
		if currInterval[0] <= newInterval[1] {
			newInterval[0] = util.Min(newInterval[0], currInterval[0])
			newInterval[1] = util.Max(newInterval[1], currInterval[1])
			pos = i + 1
		} else {
			pos = i
			break
		}
	}

	ans = append(ans, newInterval)

	for i := pos; i < n; i++ {
		ans = append(ans, intervals[i])
	}

	return ans
}
