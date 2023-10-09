package dp

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #322
TC---> exponential {very slow}
SC--->O(n) {call stack}
Using recursion only. [will give TLE]
*/
func CoinChange(coins []int, amount int) int {
	minCoin := helperRec(coins, amount)

	if minCoin == math.MaxInt {
		return -1
	}

	return minCoin
}

func helperRec(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	count := math.MaxInt

	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			ans := helperRec(coins, amount-coins[i])

			if ans != math.MaxInt {
				count = util.Min(count, 1+ans)
			}
		}
	}
	return count
}

/*
Recursion + Memoization
TC--->O(amount*len(coins))
*/
func CoinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}

	minCoin := helperMem(coins, amount, &dp)

	return minCoin
}

func helperMem(coins []int, amount int, dp *[]int) int {
	if amount == 0 {
		return 0
	}

	if (*dp)[amount] != -1 {
		return (*dp)[amount]
	}

	count := math.MaxInt

	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			ans := helperMem(coins, amount-coins[i], dp)
			if ans != math.MaxInt {
				count = util.Min(count, 1+ans)
			}
		}
	}
	(*dp)[amount] = count

	return (*dp)[amount]
}

/*
Tabulation method
TC--->O(amount*len(coins))
SC--->O(amount)
*/
func CoinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for val := 1; val < amount+1; val++ {
		for i := 0; i < len(coins); i++ {
			if amount >= coins[i] && dp[val-coins[i]] != math.MaxInt {
				dp[val] = util.Min(1+dp[val-coins[i]], dp[val])
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}
