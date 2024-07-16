package findpath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGraph(t *testing.T) {
	/** given a graph with no edges */
	n := 0
	edges := [][]int{}
	source := 0
	destination := 0
	expected := false
	result := validPath(n, edges, source, destination)
	assert.Equal(t, expected, result, "Empty graph should return false as there are no edges")
}

func TestSingleEdgeGraph(t *testing.T) {
	/** given a graph with a single edge */
	n := 2
	edges := [][]int{{0, 1}}
	source := 0
	destination := 1
	expected := true
	result := validPath(n, edges, source, destination)
	assert.Equal(t, expected, result, "Graph with a single edge should return true if source and destination are connected")
}

func TestDoubleEdgeGraphDisconnected(t *testing.T) {
	/** given a graph with a single edge */
	n := 2
	edges := [][]int{{0, 1}, {2, 3}}
	source := 2
	destination := 0
	expected := false
	result := validPath(n, edges, source, destination)
	assert.Equal(t, expected, result, "Graph with a single edge should return false if source and destination are not connected")
}

/*
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
Output: true
*/
func TestTripleEdgeGraph(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	destination := 2
	expected := true
	result := validPath(n, edges, source, destination)
	assert.Equal(t, expected, result, "Graph with a single edge should return true if source and destination are connected")
}
