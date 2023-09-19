package util

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Average(a, b int) float64 {
	return (float64(a) + float64(b)) / 2
}

// Returns 0 if a & b equal, 1 if a > b and -1 if a < b
func Signum(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
