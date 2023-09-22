package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	length := LengthOfList(head)

	pos := length - n
	dummy := &ListNode{Next: head}
	prev := dummy

	currNode := prev.Next

	for currNode != nil {
		if pos == 0 {
			break
		}

		pos--
		prev = currNode
		currNode = currNode.Next
	}

	prev.Next = currNode.Next
	currNode.Next = nil

	return dummy.Next
}

/*
Two-Pointer approach.
*/
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	node1, node2 := prev.Next, prev.Next

	for n > 0 {
		n--
		node2 = node2.Next
	}

	// if node2 is nil then n = length of linked list
	if node2 == nil {
		return head.Next
	}

	for node2.Next != nil {
		node1 = node1.Next
		node2 = node2.Next
	}

	node1.Next = node1.Next.Next

	return dummy.Next
}
