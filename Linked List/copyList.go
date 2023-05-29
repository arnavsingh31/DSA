package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyList(head *Node) *Node {
	copyHead := &Node{}
	currNode := copyHead
	originalNode := head
	originalToNewNodeMap := make(map[*Node]*Node)

	for originalNode != nil {
		copyNode := new(Node)
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
