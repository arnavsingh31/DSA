package main

func lcm(a, b int) int {
	max := large(a, b)

	for max <= a*b {
		if max%a == 0 && max%b == 0 {
			break
		}

		max++
	}
	return max
}

func large(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// a*b = gcd(a,b)*lcm(a,b)
func lcm2(a, b int) int {
	return (a * b) / hcf(a, b)
}
func hcf(a, b int) int {
	if b == 0 {
		return a
	}

	return hcf(b, a%b)
}
