package linkedlist

// LC #203
func RemoveElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	currNode := prevNode.Next

	for currNode != nil {
		if currNode.Val == val {
			prevNode.Next = currNode.Next
			currNode = currNode.Next
		} else {
			prevNode = currNode
			currNode = currNode.Next
		}
	}

	return dummyNode.Next
}
