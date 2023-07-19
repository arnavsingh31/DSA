package main

// Solution is not optimal redo it after learning stack data structure.
func nextGreater(head *ListNode) []int {
	var searchNode *ListNode
	ans := []int{}
	currNode := head

	for currNode != nil {
		searchNode = currNode.Next
		if searchNode == nil {
			ans = append(ans, 0)
		}
		for searchNode != nil {
			if searchNode.Val > currNode.Val {
				ans = append(ans, searchNode.Val)
				break
			}
			searchNode = searchNode.Next
			if searchNode == nil {
				ans = append(ans, 0)
			}
		}
		currNode = currNode.Next
	}

	return ans
}

/*
	Optimised Solution using monotonic stack
	T.C--> O(n)
	S.C--> O(n)

*/
func nextGreater2(head *ListNode) []int {
	stack := []int{}
	arr := make([]int, 0)
	nextGr := make([]int, 0)

	// convert linked list value to array
	for currNode := head; currNode != nil; currNode = currNode.Next {
		arr = append(arr, currNode.Val)
		nextGr = append(nextGr, 0)
	}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGr[stackTop] = arr[i]
		}

		stack = append(stack, i)
	}
	return nextGr
}
