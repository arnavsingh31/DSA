package bitwiseoperation

// return sum of 2 integers without using + operator
func GetSum(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}
