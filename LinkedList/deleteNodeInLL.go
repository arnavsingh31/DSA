package linkedlist

// LC #237
func DeleteNode(node *ListNode) {
	var prevNode *ListNode
	currNode := node

	for currNode.Next != nil {
		prevNode = currNode
		currNode = currNode.Next
		prevNode.Val = currNode.Val
	}
	prevNode.Next = nil
}
