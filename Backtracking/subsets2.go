package backtracking

import "sort"

/*
	LC #90
	TC--->O(2^n{recursion} * n{copying data to ans})
	SC--->O(2^n{subsets}* k{average length of a subset} )
*/
func SubsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	solveRecSub(nums, []int{}, 0, &ans)
	return ans
}

/*
	base case not needed here, since we are iterating from index to n-1 so when we come to last index, the loop
	won't execute and func call will automatically terminate.

	Algo steps :
	1. Here we’ll first sort the array and for every index, we’ll try to push elements in such a way that it doesn’t contain any duplicates.

	2. In every function call, we’ll push our temp to ans vector. For the first call, it will push an empty subset.

	3. Now we will iterate from index to n-1 & try to pick & not pick an element.

	4. To avoid duplicates if we have already picked an element for the current index of curr with the same value as the index of the given array then we’ll continue. Among duplicates we only pick its first instance(i.e when i== index) for rest other duplicate instances(i != index && it is equal to prev val) we continue

	5. After every call we’ll pop our current element.
*/
func solveRecSub(arr []int, curr []int, index int, ans *[][]int) {
	temp := append([]int{}, curr...)
	*ans = append(*ans, temp)

	for i := index; i < len(arr); i++ {
		if i != index && arr[i] == arr[i-1] {
			continue
		}

		curr = append(curr, arr[i])
		solveRecSub(arr, curr, i+1, ans)
		curr = curr[:len(curr)-1]
	}
}
