package arrays

import util "github.com/arnavsingh31/DSA/Util"

/*
LC #122
TC--->O(n)
SC--->O(1)
*/
func MaxProfit2(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for _, currPrice := range prices {
		if currPrice < minPrice {
			minPrice = currPrice
		} else {
			profit += currPrice - minPrice
			minPrice = currPrice
		}
	}
	return profit
}

/*
LC #122
TC--->O(2^n)
SC--->O(n)
*/
func MaxProfitRec(prices []int) int {
	return calcProfit(prices, 0, true)
}

func calcProfit(arr []int, index int, canBuy bool) int {
	if index == len(arr) {
		return 0
	}

	profit := 0

	if canBuy {
		profit = util.Max(-arr[index]+calcProfit(arr, index+1, false), 0+calcProfit(arr, index+1, true))
	} else {
		profit = util.Max(arr[index]+calcProfit(arr, index+1, true), 0+calcProfit(arr, index+1, false))
	}

	return profit
}

func MaxProfitDP(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}

	return calcProfitDP(prices, 0, 1, &dp)
}

func calcProfitDP(arr []int, index int, canBuy int, dp *[][]int) int {
	if index == len(arr) {
		return 0
	}

	if (*dp)[index][canBuy] != -1 {
		return (*dp)[index][canBuy]
	}

	profit := 0

	if canBuy == 1 {
		profit = util.Max(-arr[index]+calcProfitDP(arr, index+1, 0, dp), 0+calcProfitDP(arr, index+1, 1, dp))
	} else {
		profit = util.Max(arr[index]+calcProfitDP(arr, index+1, 1, dp), 0+calcProfitDP(arr, index+1, 0, dp))
	}

	(*dp)[index][canBuy] = profit

	return profit
}
