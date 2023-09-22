package linkedlist

func RemoveDuplicatesFromSortedList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: -101, Next: head}
	deleteNode := &ListNode{}
	nextDeleteNode := &ListNode{}
	prev := dummy

	currNode := prev.Next

	for currNode != nil {

		if (currNode.Next != nil) && (currNode.Val == currNode.Next.Val) {
			deleteNode = currNode
			nextDeleteNode = currNode.Next
		}

		if currNode == deleteNode || currNode == nextDeleteNode {
			prev.Next = currNode.Next
			currNode = currNode.Next
			continue
		}

		prev = currNode
		currNode = currNode.Next
	}

	return dummy.Next
}
