package main

/*
	LC #121
	TC--->O(n)
	SC--->O(1)
*/
func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, currPrice := range prices {
		if currPrice < minPrice {
			minPrice = currPrice
		} else {
			profit := currPrice - minPrice
			maxProfit = maxPrice(maxProfit, profit)
		}
	}
	return maxProfit
}

func maxPrice(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*


 */
