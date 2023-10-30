package bitwiseoperation

import (
	"strconv"
)

// Use full-adder implementation
func AddBinaryString(a, b string) string {
	var result string
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	carry := 0
	for aIndex >= 0 || bIndex >= 0 {
		x, y := 0, 0

		if aIndex >= 0 {
			x = int(a[aIndex] - '0')
			aIndex--
		}

		if bIndex >= 0 {
			y = int(b[bIndex] - '0')
			bIndex--
		}

		sum := x ^ y ^ carry
		carry = (x & y) | ((x ^ y) & carry)
		result = strconv.Itoa(sum) + result
	}

	if carry == 1 {
		result = "1" + result
	}

	return result

}
