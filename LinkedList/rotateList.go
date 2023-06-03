package main

// APPROACH #1 using recurrsion. S.C ---> O(k) ; T.C ---> O(n.k){i think so}
func rotateList(head *ListNode, k int) *ListNode {
	if head == nil || (k == lengthOfList(head)) {
		return head
	}

	k %= lengthOfList(head)

	head = recursionRotate(head, k)

	return head
}

func recursionRotate(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	currNode := prevNode.Next

	for currNode != nil {
		if currNode.Next == nil {
			prevNode.Next = nil
			currNode.Next = dummyNode.Next
			dummyNode.Next = currNode
			break
		}
		prevNode = currNode
		currNode = currNode.Next
	}

	return recursionRotate(dummyNode.Next, k-1)
}

// APPROACH #2 without using extra space .ie S.C.--> O(1) and T.C---> O(n)
func rotateList2(head *ListNode, k int) *ListNode {
	length := lengthOfList(head)
	if head == nil || k == 0 || (k == length) || (length == 1) {
		return head
	}

	k %= length

	newHeadNode := &ListNode{}
	dummyNode := &ListNode{Next: head}
	prevNode := dummyNode
	currNode := prevNode.Next
	pos := 1

	for currNode != nil && k > 0 {
		if currNode.Next == nil {
			currNode.Next = dummyNode.Next
			dummyNode.Next = newHeadNode
			break
		}

		if pos == length-k {
			newHeadNode = currNode.Next
			prevNode = currNode
			currNode.Next = nil
			currNode = newHeadNode
			pos++
			continue
		}
		prevNode = currNode
		currNode = currNode.Next
		pos++
	}
	return dummyNode.Next
}
