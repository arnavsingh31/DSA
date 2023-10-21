package graph

/*
	LC #547
	TC--->O(n+2E)
	SC--->O(n) {visited array, call stack, adjList}
*/
func NumberOfProvince(adjMat [][]int) int {
	cities := len(adjMat)
	visited := make([]bool, cities)

	bfs := func(startNode int) {
		queue := []int{startNode}
		visited[startNode] = true

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for i := 0; i < cities; i++ {
				if adjMat[node][i] == 1 && !visited[i] {
					queue = append(queue, i)
					visited[i] = true
				}
			}

		}
	}

	province := 0
	for i := 0; i < cities; i++ {
		if !visited[i] {
			bfs(i)
			province++
		}
	}

	return province
}

func NumberOfProvince2(adjMat [][]int) int {
	adjList := make(map[int][]int)
	n := len(adjMat)
	visited := make([]bool, n)

	// making adjacency list form given matrix
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && adjMat[i][j] == 1 {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	// apply dfs for a node not visited
	province := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			province++
			dfs(i, adjList, &visited)
		}
	}

	return province
}

func dfs(node int, adjList map[int][]int, visited *[]bool) {
	(*visited)[node] = true
	for _, city := range adjList[node] {
		if !(*visited)[city] {
			dfs(city, adjList, visited)
		}
	}
}
