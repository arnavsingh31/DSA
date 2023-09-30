package trees

import "container/heap"

/*
	LC #230 **Note: use morris traversal algorithm for constant SC.
	TC--->O(n)
	SC--->O(n)
	Find inorder traversal of BST and store in an array and then use heap to get kth smallest value.
*/

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	lastIndex := len(*h) - 1
	element := (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	return element
}

// using heap + inordertraversal
func KthSmallestInBST(root *Node, k int) int {
	arr := getInorderTraversal(root)
	pq := make(MaxHeap, 0)

	heap.Init(&pq)

	for i := 0; i < k; i++ {
		heap.Push(&pq, arr[i])
	}

	for i := k; i < len(arr); i++ {
		top := pq[0]

		if top > arr[i] {
			heap.Pop(&pq)
			heap.Push(&pq, arr[i])
		}
	}

	return heap.Pop(&pq).(int)
}

// Left-->Root--->Right
func getInorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	left := getInorderTraversal(root.Left)
	right := getInorderTraversal(root.Right)

	return append(append(append([]int{}, left...), root.Val), right...)
}

// using only onorder traversal
func KthSmallestInBST2(root *Node, k int) int {
	arr := getInorderTraversal(root)
	var ans int
	for i := 0; i < k; i++ {
		ans = arr[i]
	}
	return ans
}

func KthSmallestInBST3(root *Node, k int) int {
	var ans int
	getKthSmallestNode(root, &k, &ans)
	return ans
}

func getKthSmallestNode(root *Node, k *int, ans *int) {
	if root == nil {
		return
	}

	getKthSmallestNode(root.Left, k, ans)

	*k--
	if *k == 0 {
		*ans = root.Val
		return
	}

	getKthSmallestNode(root.Right, k, ans)
}
