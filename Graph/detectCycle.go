package graph

/*
	GFG
	TC--->O(n+2E)
	SC--->O(n)
	Detect cycle in undirected graph using BFS
*/
func DetectCycle(v int, adjList map[int][]int) bool {
	visited := make([]bool, v)
	bfs := func(startNode int) bool {
		queue := []Pair{{startNode, -1}}
		visited[startNode] = true

		for len(queue) > 0 {
			node := queue[0].first
			parent := queue[0].second
			queue = queue[1:]

			for _, adjNode := range adjList[node] {
				if !visited[adjNode] {
					queue = append(queue, Pair{adjNode, node})
				} else if adjNode != parent {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < v; i++ {
		if !visited[i] {
			if bfs(i) {
				return true
			}
		}
	}
	return false
}

type Pair struct {
	// first --> adjacent node, second--> parent
	first, second int
}
