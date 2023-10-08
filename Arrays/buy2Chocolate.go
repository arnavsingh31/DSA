package arrays

// approach using two pointers
func BuyChoco(prices []int, money int) int {
	left, right := 0, len(prices)-1

	savings := -1

	for left < right {
		cost := prices[left] + prices[right]
		if money-cost > savings {
			savings = money - cost
		}

		if prices[left] > prices[right] {
			left++
		} else {
			right--
		}
	}

	if savings < 0 {
		return money
	}
	return savings
}
