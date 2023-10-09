package dp

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #746
	TC--->O(n)
	SC--->O(n)
	Recursion {TLE}
*/
func MinCostClimbingStairs(cost []int) int {
	top := len(cost)
	return util.Min(solveRec1(cost, top-1), solveRec1(cost, top-2))
}

func solveRec1(cost []int, pos int) int {
	if pos < 0 {
		return 0
	}
	if pos == 0 || pos == 1 {
		return cost[pos]
	}

	ans := util.Min(solveRec1(cost, pos-1), solveRec1(cost, pos-2))
	return ans + cost[pos]
}

/*
	Recursion + Memoization
	TC--->O(n)
	SC--->O(n)+O(n)~ O(n) {call stack + dp array}
*/
func MinCostClimbingStairs2(cost []int) int {
	top := len(cost)
	dp := make([]int, len(cost)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	return util.Min(solveMem(cost, top-1, &dp), solveMem(cost, top-2, &dp))
}

func solveMem(cost []int, pos int, dp *[]int) int {
	if pos < 0 {
		return 0
	}

	if pos == 0 || pos == 1 {
		return cost[pos]
	}

	if (*dp)[pos] != -1 {
		return (*dp)[pos]
	}

	ans := util.Min(solveMem(cost, pos-1, dp), solveMem(cost, pos-2, dp))
	(*dp)[pos] = ans + cost[pos]

	return (*dp)[pos]
}

/*

	Tabulation Method (though SC is same for both but uses less space{only dp} than memoization solution{dp+ call stack} )
	TC--->O(n)
	SC--->O(n) {only dp array}
*/

func MinCostClimbingStairs3(cost []int) int {
	top := len(cost)
	dp := make([]int, top+1)
	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < top; i++ {
		dp[i] = cost[i] + util.Min(dp[i-1], dp[i-2])
	}

	return util.Min(dp[top-1], dp[top-2])
}

/*
	Tabulation with space optimisation.
	TC--->O(n)
	SC--->O(1)
*/
func MinCostClimbingStairs4(cost []int) int {
	top := len(cost)
	prev1 := cost[1]
	prev2 := cost[0]

	for i := 2; i < top; i++ {
		curr := cost[i] + util.Min(prev1, prev2)
		prev2 = prev1
		prev1 = curr
	}

	return util.Min(prev1, prev2)
}
