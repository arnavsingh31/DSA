package maths

// Sieve of Eratosthenes algorithm
// T.C ---> O(n*log(logn)) (proof is too complex)
// we can avoid extra space by directly printing the prime numbers.
func PrimeNumbers(n int) []int {
	arr := []bool{true, true}
	for i := 2; i <= n+1; i++ {
		arr = append(arr, true)
	}

	for i := 2; i*i <= n; i++ {
		if arr[i] {
			for j := 2 * i; j <= n; j += i {
				arr[j] = false
			}
		}
	}
	ans := []int{}
	for i := 1; i <= n; i++ {
		if arr[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

// T.C ---> O(n*log(logn)) (proof is too complex)
// we can avoid extra space by directly printing the prime numbers.
func PrimeNumbers2(n int) []int {
	arr := []bool{true, true}
	for i := 2; i <= n+1; i++ {
		arr = append(arr, true)
	}
	ans := []int{}
	for i := 2; i <= n; i++ {
		if arr[i] {
			ans = append(ans, i)
			for j := i * i; j <= n; j += i {
				arr[j] = false
			}
		}
	}
	return ans
}
