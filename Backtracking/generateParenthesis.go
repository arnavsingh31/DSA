package backtracking

/*
	LC #22
	TC--->O((2^n)*n)
	SC--->O(2*N*X), X->number of valid parenthesis
*/
func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)
	solveRecP(n, 0, 0, "", &ans)
	return ans
}

func solveRecP(n int, open, close int, curr string, ans *[]string) {
	if len(curr) == n*2 {
		*ans = append(*ans, curr)
		return
	}

	if open < n {
		solveRecP(n, open+1, close, curr+"(", ans)
	}

	if close < open {
		solveRecP(n, open, close+1, curr+")", ans)
	}
}
