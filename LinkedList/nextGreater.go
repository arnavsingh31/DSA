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
