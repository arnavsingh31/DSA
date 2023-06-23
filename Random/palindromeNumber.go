package main

func isNumberPalindrome(num int) bool {
	rev := 0
	x := num

	for num > 0 {
		lastDigit := num % 10
		rev = rev*10 + lastDigit
		num = num / 10
	}

	return rev == x
}
