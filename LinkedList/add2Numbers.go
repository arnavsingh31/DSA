package main

func add2Numbers(list1, list2 *ListNode) *ListNode {
	var ansHead, ansPrev *ListNode

	curr1 := revLinkedList(list1)
	curr2 := revLinkedList(list2)
	sum, carry := 0, 0

	for curr1 != nil || curr2 != nil {
		if curr1 != nil {
			sum += curr1.Val
			curr1 = curr1.Next
		}

		if curr2 != nil {
			sum += curr2.Val
			curr2 = curr2.Next
		}

		ansNode := &ListNode{Val: sum % 10}
		ansHead = ansNode
		carry = sum / 10
		ansNode.Next = ansPrev
		ansPrev = ansNode
		sum = carry
	}

	if carry != 0 {
		newNode := &ListNode{Val: carry, Next: ansHead}
		return newNode
	}

	return ansHead
}

func revLinkedList(head *ListNode) *ListNode {
	var prevNode, tempNode *ListNode
	currNode := head

	for currNode != nil {
		tempNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
	}

	return prevNode
}
