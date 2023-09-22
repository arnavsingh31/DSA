package linkedlist

func PartitionList(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	largeHead := &ListNode{}
	largePtr := &ListNode{}
	smallPtr := &ListNode{}
	smallhead := &ListNode{}

	currNode := dummyNode.Next

	for currNode != nil {
		if currNode.Val < x {
			if smallhead.Next == nil {
				smallhead.Next = currNode
			}
			smallPtr.Next = currNode
			smallPtr = currNode
		} else { // currNode.Val >= x
			if largeHead.Next == nil {
				largeHead.Next = currNode
			}
			largePtr.Next = currNode
			largePtr = currNode
		}
		currNode = currNode.Next
	}
	if smallhead.Next == nil {
		return largeHead.Next
	}
	smallPtr.Next = largeHead.Next
	largePtr.Next = nil

	return smallhead.Next
}
