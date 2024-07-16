package simplebfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleBfs(t *testing.T) {
	// given the graph:
	// 0 - 1 - 2
	// |   |   |
	// 3 - 4 - 5
	// the adjacency list is:
	// [[1, 3], [0, 2, 4], [1, 5], [0, 4], [1, 3, 5], [2, 4]]
	adjList := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

	assert.Equal(t, []int{0, 1, 3, 2, 4, 5}, traverse(adjList, 0), "Should return the correct traversal for the given graph")
}
