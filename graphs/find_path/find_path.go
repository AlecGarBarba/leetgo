package findpath

func validPath(n int, edges [][]int, source int, destination int) bool {

	if len(edges) == 0 && source != destination {
		return false
	}

	if source == destination {
		return true
	}

	adjList := [][]int{}

	/** initialize the adjacency list */
	for i := 0; i < n; i++ {
		adjList = append(adjList, []int{})
	}

	/** populate the adjacency list */
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	visited := make([]bool, len(adjList))

	visited[source] = true
	q := []int{source}

	for len(q) > 0 {
		curr := q[0]
		// dequeue
		q = q[1:]

		/*
			We need to iterate over the neighbors of the current node and add them to the queue if they haven't been visited yet.
		*/
		for _, neighbor := range adjList[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				// enqueue
				q = append(q, neighbor)
			}
		}
	}

	return visited[destination]

}
