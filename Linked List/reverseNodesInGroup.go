package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// APPROACH #1 Recursion.
// We will reverse first k nodes rest recursion will handle.
func reverseKGroup(head *ListNode, k int) *ListNode {

	intitialDiff := lengthOfList(head)

	head = helper(head, k, intitialDiff)

	return head
}

func lengthOfList(head *ListNode) int {
	if head == nil {
		return 0
	}

	len := 0
	currNode := head
	for currNode != nil {
		currNode = currNode.Next
		len++
	}
	return len
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
