package bitwiseoperation

func IsPowOf2(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

func IsPowerOf2(n int) bool {
	if n == 0 {
		return false
	}
	n = n & (n - 1)
	return n == 0
}
