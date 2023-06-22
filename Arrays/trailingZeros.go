package main

func trailingZeros(n int) int {
	count_two := 0
	count_five := 0

	for i := 2; i <= n; i++ {
		temp := i

		for temp > 1 {
			if temp%2 == 0 {
				count_two++
				temp = temp / 2

			} else if temp%5 == 0 {
				count_five++
				temp = temp / 5

			} else {
				break
			}
		}
	}

	return min(count_two, count_five)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
