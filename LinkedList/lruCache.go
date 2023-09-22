package linkedlist

/*
Approach
- We use a doubly-linked list and hashmap to perform this operation.
- head of linked list denotes the latest recently used item and tail of the list denotes the least recently used item in the list.
- hashmap is a map of key to its corresponding node in the linked list.

Complexity
- Time complexity:
	O(1) for both get and put.

	For get():

	- Check if a key is in a hash map. This costs O(1).
	- Get a node associated with a key. This costs O(1).
	- Then calls the MoveNode() method. This costs O(1).

	For put():

	- Check if a key is in a hash map. This costs O(1).
	- If it exist then we update the val in the existing node. This costs O(1).
	- Then we move the node back in front of the list, as this is latest used node. This costs O(1).
	- If key doesn't exist and capacity of cache is reached, then we simply delete the key from our map. This costs O(1).
	- After deleting key from map we also delete corresponding node which is LRU node i.e. tail node and point the tail to node prev to the current tail. This costs O(1).
	- After this we create new node and call our addNode method. This costs again O(1).

- Space complexity:
	We use a hashmap and a doubly-linked list which use extra space at most equal to the size of capacity. Hence we get space complexity of O(capacity).
*/

type Cache map[int]*DoubleListNode

type DoubleListNode struct {
	Key  int
	Val  int
	Prev *DoubleListNode
	Next *DoubleListNode
}

type LRUCache struct {
	head     *DoubleListNode
	tail     *DoubleListNode
	cache    Cache
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     &DoubleListNode{},
		tail:     &DoubleListNode{},
		cache:    make(Cache),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	node, exist := c.cache[key]

	if exist {
		// modify the lru linked list by moving the node to front again.
		c.MoveNode(node)
		return node.Val
	}

	return -1
}

func (c *LRUCache) MoveNode(node *DoubleListNode) {
	if node == c.tail.Prev {
		c.tail.Prev = node.Prev
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = c.head.Next
	node.Prev = c.head
	c.head.Next.Prev = node
	c.head.Next = node
}

func (c *LRUCache) Put(key, val int) {
	mapLen := len(c.cache)

	node, exist := c.cache[key]
	if exist {
		node.Val = val
		// move this node back in front as this is the latest used node in linked list
		c.MoveNode(node)
		return
	}

	if mapLen == c.capacity && !exist {
		// delete least recently used item from list as well as from map
		// delete lru item from map
		delete(c.cache, c.tail.Prev.Key)

		// delete lru node from linked list which is the tail node and modify tail
		tail := c.tail.Prev
		tail.Prev.Next = tail.Next
		c.tail.Prev = tail.Prev
	}

	// now we can add new key to the map/cache/list
	newNode := &DoubleListNode{Key: key, Val: val}
	c.cache[key] = newNode
	c.addNode(newNode)
}

func (c *LRUCache) addNode(node *DoubleListNode) {
	if c.head.Next == nil {
		c.head.Next = node
		c.tail.Prev = node
		node.Next = c.tail
		node.Prev = c.head
		return
	}

	node.Next = c.head.Next
	node.Prev = c.head
	c.head.Next = node
}
