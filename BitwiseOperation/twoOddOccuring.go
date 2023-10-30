package bitwiseoperation

// T.C---> theta(n)
func TwoOddOccuring(arr []int) []int {
	res := 0

	for _, num := range arr {
		res = res ^ num
	}
	/*
		how does res & ^(res-1) work?
		it finds a number which has only 1 bit set and the set bit corresponds to last bit of x
		after finding the number(temp), we find all those number in array whose corresponding bit is also
		set and do XOR of all such numbers which will give us one of the two odd occuring numbers.
		two get another odd occuring number we do XOR of all numbers whose corresponding bit is not set.
	*/
	temp := res & (^(res - 1))
	x, y := 0, 0
	for _, num := range arr {
		if temp&num != 0 {
			x = x ^ num
		} else {
			y = y ^ num

		}
	}

	return []int{x, y}
}
