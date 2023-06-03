package main

func reverseBetween(head *ListNode, left, right int) *ListNode {
	if head == nil {
		return head
	}

	var prevNode, tempNode, beforeLeftNode *ListNode
	currNode := head
	pos := 1
	dummy := &ListNode{Val: 0, Next: head}
	beforeLeftNode = dummy
	for pos < left {
		beforeLeftNode = currNode
		currNode = currNode.Next
		pos++
	}

	for pos <= right {
		tempNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
		pos++
	}

	beforeLeftNode.Next.Next = currNode
	beforeLeftNode.Next = prevNode

	return dummy.Next
}
