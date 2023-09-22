package linkedlist

func CycleStart(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	currentNode := head
	nodeVisitedMap := make(map[*ListNode]bool)
	for currentNode != nil {
		if nodeVisitedMap[currentNode] {
			return currentNode
		}

		nodeVisitedMap[currentNode] = true
		currentNode = currentNode.Next
	}

	return nil
}

// if cycle is present reset fast to head and keep slow at it current node and now move slow and fast at
// same pace and where they will meet is the starting node of cycle.
func CycleStart2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
