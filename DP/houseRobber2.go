package main

/*
	LC #213
	TC--->O(2^n exponential) {TLE}, recursive sol.
	SC--->O(n)
*/
func robHouseCircular(arr []int) int {
	rob1 := arr[0] + robRec(arr, 2, true)
	dontRob1 := robRec(arr, 1, false)
	return max(rob1, dontRob1)
}

func robRec(arr []int, i int, robbedFirst bool) int {
	if i >= len(arr) {
		return 0
	}

	// if robbing 1 st house then you cannot rob last house or if robbing last house then you cannot rob first.
	if i == len(arr)-1 && robbedFirst {
		return 0
	}

	rob := arr[i] + robRec(arr, i+2, robbedFirst)
	dontRob := robRec(arr, i+1, robbedFirst)

	return max(rob, dontRob)
}

/*
	Recursion + Memoization
	TC--->O(n)
	SC--->O(2n)
*/
func robHouseCircular2(arr []int) int {
	dp := make([][]int, len(arr))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}

	rob1 := arr[0] + robHouseMem(arr, 2, 1, &dp)
	dontRob1 := robHouseMem(arr, 1, 0, &dp)

	return max(rob1, dontRob1)
}

func robHouseMem(arr []int, i int, robbedFirst int, dp *[][]int) int {
	if i >= len(arr) {
		return 0
	}

	if i == len(arr)-1 && robbedFirst == 1 {
		return 0
	}

	if (*dp)[i][robbedFirst] != -1 {
		return (*dp)[i][robbedFirst]
	}

	rob := arr[i] + robHouseMem(arr, i+2, robbedFirst, dp)
	dontRob := robHouseMem(arr, i+1, robbedFirst, dp)

	(*dp)[i][robbedFirst] = max(rob, dontRob)

	return (*dp)[i][robbedFirst]
}

/*
	TC--->O(n)
	SC--->O(1)
	We used house robber 1 solution twice.
*/
func robHouseCircular3(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	// if we rob first house then we remove last house from array
	robFirst := solveHouseRobber1(arr[:len(arr)-1])

	//if we rob last house then we remove first house from array
	robLast := solveHouseRobber1(arr[1:])

	return max(robFirst, robLast)
}

func solveHouseRobber1(arr []int) int {
	prev2, prev1 := 0, 0

	for i := 0; i < len(arr); i++ {
		rob := arr[i] + prev2
		dontRob := prev1

		prev2 = prev1
		prev1 = max(rob, dontRob)
	}

	return prev1
}
