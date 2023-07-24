package main

/*
	LC #769
	Here we used the fact that numbers in array are between 0-(n-1) so are index. we store a num in max n check
	that whether that max is at its right index or not if not we continue if yes then we have found our first
	chunk.
	TC-->O(n), SC-->O(1)
*/
func maxChunks(arr []int) int {
	chunks := 0
	max := 0

	for i, num := range arr {
		if num > max {
			max = num
		}
		if max == i {
			chunks++
		}
	}
	return chunks
}
