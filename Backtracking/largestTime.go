package backtracking

import (
	"fmt"
	"strings"
)

/*
LC #949
As per question size of array is always = 4, hence there will be only 24 permutations possible.
*/
func GetMaxTime(arr []int) string {
	ans := ""
	backtrack(&arr, 0, &ans)
	return ans
}

func backtrack(arr *[]int, start int, ans *string) {
	if start == len(*arr) {
		// check whether the existing permutation in array is valid or not
		if (*arr)[0] != 0 && (*arr)[0] != 1 && (*arr)[0] != 2 {
			return
		}

		if (*arr)[0] == 2 && (*arr)[1] > 3 {
			return
		}

		if (*arr)[2] > 5 {
			return
		}

		// now our permutation in array is valid so add it to our possible ans
		temp := fmt.Sprintf("%d%d:%d%d", (*arr)[0], (*arr)[1], (*arr)[2], (*arr)[3])

		// -1 mean temp time is > prev time
		if strings.Compare(*ans, temp) == -1 {
			*ans = temp
		}

		return
	}

	for i := start; i < len(*arr); i++ {
		// swap
		(*arr)[start], (*arr)[i] = (*arr)[i], (*arr)[start]
		backtrack(arr, start+1, ans)
		(*arr)[start], (*arr)[i] = (*arr)[i], (*arr)[start]
	}

}
