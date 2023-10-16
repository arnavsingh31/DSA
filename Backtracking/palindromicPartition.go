package backtracking

/*
	LC #131
	TC--->O(n* 2^n)
	SC--->O(2^n)

	Should be O(n*2^n). You are basically trying out every possible partition out there. For a string with length n, you will have 2^(n - 1) ways to partition it. This is because, a partition is equivalent of putting a "|" in b/t two chars. There are n - 1 such slots to place a "|". There are only two choice for each slot - placing a "|" or not placing a "|". Thus 2^(n - 1) ways to place "|"s.

	Then for each unique partitioning, you have to traverse the entire string (in the worst case when you have repeating chars) to make sure every partition is a palindrome. so n * 2 ^ (n - 1) = O(n*2^n).
*/
func PalindromicPartitions(s string) [][]string {
	ans := make([][]string, 0)
	solvePartition(s, 0, []string{}, &ans)
	return ans
}

func solvePartition(s string, index int, path []string, ans *[][]string) {
	if index >= len(s) {
		temp := append([]string{}, path...)
		*ans = append(*ans, temp)
		return
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s, index, i) {
			subStr := s[index : i+1]
			path = append(path, subStr)
			solvePartition(s, i+1, path, ans)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
