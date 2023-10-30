package bitwiseoperation

func MissingNumber(arr []int) int {
	n := len(arr)
	sum := (n * (n + 1)) / 2
	hasZero := false
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			hasZero = true
		}
		sum -= arr[i]

	}
	if sum == 0 && hasZero {
		return n
	} else if sum == 0 && !hasZero {
		return 0
	}
	return sum
}

// using bit-operation
func MissingNumber2(arr []int) int {
	sum := 0

	for i := 0; i <= len(arr); i++ {
		sum ^= i
	}

	for _, num := range arr {
		sum ^= num
	}
	return sum
}
