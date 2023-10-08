package arrays

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
