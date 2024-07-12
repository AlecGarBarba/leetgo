package centerstar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstStar(t *testing.T) {
	// given edges = [[1,2],[2,3],[4,2]]
	// the graph looks like this:
	// 1 - 2 - 3
	//     |
	//     4
	// the center is 2
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	assert.Equal(t, 2, findCenter(edges), "Should return 2")
}

func TestSecondStar(t *testing.T) {
	// given Input: edges = [[1,2],[5,1],[1,3],[1,4]]
	// Output: 1
	// the graph looks like this:
	// 2 - 1 - 3
	//   / |
	//  5  4
	// the center is 1
	edges := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
	assert.Equal(t, 1, findCenter(edges), "Should return 1")

}
