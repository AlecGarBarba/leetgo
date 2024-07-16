package keysrooms

func canVisitAllRooms(rooms [][]int) bool {

	// cover base cases

	if len(rooms) == 0 {
		return true
	}

	visited := make([]bool, len(rooms))
	visited[0] = true

	q := []int{0}

	for len(q) > 0 {
		curr := q[0]

		q = q[1:]

		for _, roomKey := range rooms[curr] {

			if !visited[roomKey] {
				visited[roomKey] = true

				q = append(q, roomKey)
			}
		}
	}

	for _, vis := range visited {
		if !vis {
			return false
		}
	}
	return true

}
