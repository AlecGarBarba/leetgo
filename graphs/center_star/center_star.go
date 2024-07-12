package centerstar

func findCenter(edges [][]int) int {

	set := make(map[int]bool)

	for _, r := range edges {

		if _, exists := set[r[0]]; exists {
			return r[0]
		}
		if _, exists := set[r[1]]; exists {
			return r[1]
		}

		set[r[0]] = true
		set[r[1]] = true
	}

	return -1

}
