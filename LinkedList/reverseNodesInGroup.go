package linkedlist

// APPROACH #1 Recursion.
// T.C---> O(n) as each node is visited only once.
// S.C---> O(n/k) as in each recursive call we reverse k nodes so to reverse 'n' nodes call stack will use space of n/k.
// We will reverse first k nodes rest recursion will handle.
func ReverseKGroup(head *ListNode, k int) *ListNode {

	intitialDiff := LengthOfList(head)

	head = helper(head, k, intitialDiff)

	return head
}

func helper(head *ListNode, k, diff int) *ListNode {
	if head == nil {
		return head
	}
	var tempNode, prevNode *ListNode
	currNode := head
	count := 0

	// step 1 reverse first k nodes
	for currNode != nil && count < k && diff >= k {
		tempNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
		count++
		if count == k {
			diff -= k
		}
	}

	// step 2
	if tempNode != nil && diff >= k {
		head.Next = helper(tempNode, k, diff)
	}

	if tempNode != nil && diff < k {
		head.Next = tempNode
	}

	// return head of reversed list
	return prevNode
}

// APPROACH #2
// T.C---> O(n)
// S.C---> O(1)
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	dummyNode := &ListNode{Next: head}
	nodeBeforeGroup := dummyNode
	tempNode := &ListNode{}
	prevNode := &ListNode{}

	for {
		kthNode := geKthNode(nodeBeforeGroup, k)

		if kthNode == nil {
			break
		}
		nodeAfterGroup := kthNode.Next

		// line 90 understand it via dry run
		prevNode = kthNode.Next
		currNode := nodeBeforeGroup.Next
		for currNode != nodeAfterGroup {
			tempNode = currNode.Next
			currNode.Next = prevNode
			prevNode = currNode
			currNode = tempNode
		}

		// line 100- ***very important. do a dry run then you will understand.
		tmp := nodeBeforeGroup.Next
		nodeBeforeGroup.Next = kthNode
		nodeBeforeGroup = tmp
	}
	return dummyNode.Next
}

func geKthNode(currNode *ListNode, k int) *ListNode {

	for currNode != nil && k > 0 {
		currNode = currNode.Next
		k--
	}
	return currNode
}
