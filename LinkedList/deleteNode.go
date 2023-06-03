package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	length := lengthOfList(head)
	nodeBeforeNthNode := length - n

	dummy := &ListNode{Next: head}
	prev := dummy

	currNode := prev.Next

	for nodeBeforeNthNode > 1 {
		currNode = currNode.Next
		nodeBeforeNthNode--
	}

	if length == n {
		prev.Next = currNode.Next
	} else {
		currNode.Next = currNode.Next.Next
	}

	return dummy.Next
}
