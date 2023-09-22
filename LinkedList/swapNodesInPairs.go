package linkedlist

// LC #24
func SwapNodesInPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	var tempNode, groupPrevNode *ListNode
	currNode := prevNode.Next
	count := 0
	for currNode != nil {
		count++
		if count%2 == 0 {
			tempNode = currNode.Next
			prevNode.Next = tempNode
			currNode.Next = prevNode
			groupPrevNode.Next = currNode
			currNode = tempNode
		} else {
			groupPrevNode = prevNode
			prevNode = currNode
			currNode = currNode.Next
		}
	}

	return dummyNode.Next
}
