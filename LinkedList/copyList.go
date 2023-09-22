package linkedlist

func CopyList(head *SpecialNode) *SpecialNode {
	copyHead := &SpecialNode{}
	currNode := copyHead
	originalNode := head
	originalToNewNodeMap := make(map[*SpecialNode]*SpecialNode)

	for originalNode != nil {
		copyNode := new(SpecialNode)
		copyNode.Val = originalNode.Val
		currNode.Next = copyNode
		currNode = copyNode
		originalToNewNodeMap[originalNode] = copyNode
		originalNode = originalNode.Next
	}

	originalNode = head

	for originalNode != nil {
		originalToNewNodeMap[originalNode].Random = originalToNewNodeMap[originalNode.Random]
		originalNode = originalNode.Next
	}

	return copyHead.Next

}
