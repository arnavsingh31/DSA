package arrays

import util "github.com/arnavsingh31/DSA/Util"

type Interval struct {
	first  int
	second int
}

// watch erichto video on YT.
func JumpGame2(arr []int) (jump int) {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	interval := &Interval{first: 0, second: 0}
	jump = 0

	for {
		jump++
		can_reach := -1

		for i := interval.first; i < interval.second; i++ {
			can_reach = util.Max(can_reach, i+arr[i])
		}

		if can_reach >= n-1 {
			return jump
		}

		//update interval
		// if inter now is [5, 8] and can_reach is 8. So after updating our intervals we cannot move any further aka we are stuck in this interval.
		interval.first, interval.second = interval.second+1, can_reach
		if interval.first > interval.second {
			return can_reach // we are stuck
		}
	}
}
