package linkedlist

/*
	we make number from value in LL nodes twice, first we make num by simply iterating the LL and
	while iterating we also reverse the LL. Then we make another reverse_num by iterating from start
	of reveresed LL. If both numbers are same, return true else return false.
*/
func IsPalindrome(head *ListNode) bool {
	var prevNode, tempNode *ListNode
	num, reverse_num := 0, 0
	currNode := head

	for currNode != nil {
		num = num*10 + currNode.Val
		tempNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tempNode
	}

	head = prevNode
	currNode = head
	for currNode != nil {
		reverse_num = 10*reverse_num + currNode.Val
		currNode = currNode.Next
	}

	return num == reverse_num
}

/*
	Using fast and slow algorithm we find end of first half of LL and the we reverse the second half of LL.
	Iterate from start to end while second half is not nil ** This is because incase of odd number of nodes
	in LL there will be an extra node in the first half which is gonna be same in plaindrome, so we iterate
	till second half is not nil. While iterating we compare node value of firstHalf and second half of LL.
	If at any point value are not equal we return false. Else at the end we return true.
*/
func IsPalindrome2(head *ListNode) bool {
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var tempNode, prevNode *ListNode
	secondHalfCurrNode := slow.Next
	for secondHalfCurrNode != nil {
		tempNode = secondHalfCurrNode.Next
		secondHalfCurrNode.Next = prevNode
		prevNode = secondHalfCurrNode
		secondHalfCurrNode = tempNode
	}
	secondHalfCurrNode = prevNode
	firstHalfCurrNode := head
	for secondHalfCurrNode != nil {
		if secondHalfCurrNode.Val != firstHalfCurrNode.Val {
			return false
		}

		secondHalfCurrNode = secondHalfCurrNode.Next
		firstHalfCurrNode = firstHalfCurrNode.Next
	}

	return true

}
