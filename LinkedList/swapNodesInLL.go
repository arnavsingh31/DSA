package linkedlist

// LC #1721
func SwapNodes(head *ListNode, k int) *ListNode {
	var nodeFromBeg, nodeFromEnd, prev1, prev2 *ListNode
	dummyNode := &ListNode{Next: head}
	prev := dummyNode
	currNode := prev.Next
	length := LengthOfList(head)
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
