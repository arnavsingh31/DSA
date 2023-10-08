package arrays

// Brute force approach T.C = O(n^2) S.C= O(n). Get TLE for large input.
func ProductArray(arr []int) []int {

	res := []int{}

	for i := 0; i < len(arr); i++ {
		product := 1
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			product *= arr[j]
		}

		res = append(res, product)
	}
	return res
}

// [7,2,3,4]
// Approach create prefix_array_product and post_array_product and simply multiply both for each i in arr
// You can make two separate array for prefix and postfix but then our S.C will be O(n).
func ProductArray2(arr []int) []int {
	res := make([]int, len(arr))

	temp := 1
	res[0] = 1 // first element doesn't have any prefix so we add 1.
	for i := 1; i < len(arr); i++ {
		res[i] = arr[i-1] * temp
		temp = res[i]
	}

	temp = 1
	for i := len(arr) - 1; i >= 0; i-- {
		// in res array we have our prefix product to which we multiply temp which is postfix_product for that ith position.
		res[i] = res[i] * temp
		temp = temp * arr[i]
	}

	return res
}
