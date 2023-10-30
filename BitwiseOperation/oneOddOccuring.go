package bitwiseoperation

// x ^ x ^ x....x (x occurs odd times) = x
// x ^ x ^ x....x (x occurs even times) = 0
// T.C--> theta(n)
func OneOddOccuring(arr []int) int {

	res := 0
	for i := 0; i < len(arr); i++ {
		res = arr[i] ^ res
	}
	return res
}
