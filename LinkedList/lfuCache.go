package main

type DoubleLinkedListNode struct {
	Key       int
	Val       int
	Frequency int
	Next      *DoubleLinkedListNode
	Prev      *DoubleLinkedListNode
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
	size int
}

type Cache1 map[int]*DoubleLinkedList
type Cache2 map[int]*DoubleLinkedListNode

type LFUCache struct {
	freqCache Cache1
	nodeCache Cache2
	capacity  int
	minFreq   int
}

func LFUConstructor(capacity int) *LFUCache {
	return &LFUCache{
		freqCache: make(Cache1),
		nodeCache: make(Cache2),
		capacity:  capacity,
		minFreq:   0,
	}
}

func MakeDoubleLinkedList() *DoubleLinkedList {
	head := &DoubleLinkedListNode{}
	tail := &DoubleLinkedListNode{}
	head.Next = tail
	tail.Prev = head

	return &DoubleLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (c *LFUCache) Get(key int) int {
	node, nodeExist := c.nodeCache[key]
	if !nodeExist {
		return -1
	}

	oldList, oldListExist := c.freqCache[node.Frequency]
	if oldListExist {
		oldList.DeleteNode(node)
	}

	node.Frequency++
	nextList, nextListExist := c.freqCache[node.Frequency]

	if !nextListExist {
		nextList = MakeDoubleLinkedList()
	}
	nextList.AddNode(node)
	c.freqCache[node.Frequency] = nextList

	if oldList.size == 0 && c.minFreq == node.Frequency-1 {
		c.minFreq++
	}

	return node.Val
}

func (list *DoubleLinkedList) DeleteNode(node *DoubleLinkedListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.size--
}

func (list *DoubleLinkedList) AddNode(node *DoubleLinkedListNode) {
	node.Next = list.head.Next
	list.head.Next.Prev = node
	list.head.Next = node
	node.Prev = list.head
	list.size++
}

func (c *LFUCache) Put(key int, value int) {
	node, nodeExist := c.nodeCache[key]

	if nodeExist {
		node.Val = value
		c.Get(key)
		return
	}

	if len(c.nodeCache) == c.capacity {
		list := c.freqCache[c.minFreq]
		leastFrequencyNode := list.tail.Prev
		list.DeleteNode(leastFrequencyNode)
		delete(c.nodeCache, leastFrequencyNode.Key)
	}

	newNode := &DoubleLinkedListNode{
		Key:       key,
		Val:       value,
		Frequency: 1,
	}

	c.minFreq = 1
	list, listExist := c.freqCache[newNode.Frequency]
	if !listExist {
		list = MakeDoubleLinkedList()
	}
	list.AddNode(newNode)
	c.freqCache[newNode.Frequency] = list
	c.nodeCache[key] = newNode
}
