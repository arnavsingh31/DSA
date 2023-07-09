package main

// LC #1721
func swapNodes(head *ListNode, k int) *ListNode {
	var nodeFromBeg, nodeFromEnd, prev1, prev2 *ListNode
	dummyNode := &ListNode{Next: head}
	prev := dummyNode
	currNode := prev.Next
	length := lenOfList(head)
	count := 0
	for currNode != nil {
		count++

		if count == k {
			prev1 = prev
			nodeFromBeg = currNode
		}

		if count == length-k+1 {
			prev2 = prev
			nodeFromEnd = currNode
		}
		prev = currNode
		currNode = currNode.Next
	}

	prev1.Next, prev2.Next = nodeFromEnd, nodeFromBeg
	nodeFromBeg.Next, nodeFromEnd.Next = nodeFromEnd.Next, nodeFromBeg.Next

	return dummyNode.Next
}

func lenOfList(head *ListNode) int {
	if head == nil {
		return 0
	}
	len := 0
	currNode := head

	for currNode != nil {
		len++
		currNode = currNode.Next
	}

	return len
}
