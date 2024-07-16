package simplebfs

func traverse(adjList [][]int, startNode int) []int {

	q := []int{startNode}

	res := []int{}
	visited := make([]bool, len(adjList))
	visited[0] = true

	for len(q) > 0 {

		curr := q[0]
		// dequeue
		q = q[1:]

		res = append(res, curr)
		for _, neighbor := range adjList[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				// enqueue
				q = append(q, neighbor)
			}
		}
	}

	return res

}
