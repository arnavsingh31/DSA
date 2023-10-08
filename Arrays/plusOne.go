package arrays

func PlusOne(digits []int) []int {
	n := len(digits)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
