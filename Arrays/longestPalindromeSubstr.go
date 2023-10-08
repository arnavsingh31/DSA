package arrays

/*
	LC #5
	TC--->O(n^2)
	SC--->O(1)
*/
func LongestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	ans := ""
	for i := 0; i < len(s); i++ {
		oddStr := expand(i, i, s)
		if len(ans) < len(oddStr) {
			ans = oddStr
		}

		evenStr := expand(i, i+1, s)
		if len(ans) < len(evenStr) {
			ans = evenStr
		}
	}

	return ans
}

func expand(left, right int, s string) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
