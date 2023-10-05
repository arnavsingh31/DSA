package maths

import util "github.com/arnavsingh31/DSA/Util"

func Lcm(a, b int) int {
	max := util.Max(a, b)

	for max <= a*b {
		if max%a == 0 && max%b == 0 {
			break
		}

		max++
	}
	return max
}

// a*b = gcd(a,b)*lcm(a,b)
func Lcm2(a, b int) int {
	return (a * b) / Gcd3(a, b)
}
