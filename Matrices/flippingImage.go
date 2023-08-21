package main

/*
	LC #832
	TC--->O(m*n)
	SC--->O(1)
*/
func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			// invert
			image[i][j] ^= 1
			image[i][n-1-j] ^= 1

			// swap
			image[i][j], image[i][n-1-j] = image[i][n-1-j], image[i][j]
		}

		if n%2 != 0 {
			image[i][n/2] ^= 1
		}
	}
	return image
}
