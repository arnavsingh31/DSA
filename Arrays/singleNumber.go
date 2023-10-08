package arrays

// Approach bit-manipulation. Doing XOR of all terms. since there are two copies of every int except one
// since xor operation of same number results in zero hence all duplicates will cancel each other ...only
// single number will exist.
func SingleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res = res ^ num
	}
	return res
}
