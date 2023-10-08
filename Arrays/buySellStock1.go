package arrays

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #121
	TC--->O(n)
	SC--->O(1)
*/
func MaxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, currPrice := range prices {
		if currPrice < minPrice {
			minPrice = currPrice
		} else {
			profit := currPrice - minPrice
			maxProfit = util.Max(maxProfit, profit)
		}
	}
	return maxProfit
}
