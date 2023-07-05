package main

/*
	1 -> 2 -> 3 -> 4 -> 5 -> 6
*/

func reverse(head *ListNode) *ListNode {
	var prevNode *ListNode
	currNode := head

	for currNode != nil {
		tempNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
	}
	head = currNode

	return head
}
