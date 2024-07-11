package townjudge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyAdjacencyList(t *testing.T) {
	// input: n = 1, trust = []
	// output: 1
	n := 1
	trust := [][]int{}
	expected := 1
	assert.Equal(t, expected, findJudge(n, trust), "Judge id should be 1")
}

func TestSinglenode(t *testing.T) {
	// input: n = 2, trust = [[1,2]]
	// output: 2
	n := 2
	trust := [][]int{{1, 2}}
	expected := 2
	assert.Equal(t, expected, findJudge(n, trust), "Judge id should be 2")

}

func TestTwoNodes(t *testing.T) {
	// input: n = 3, trust = [[1,3],[2,3]]
	// output: 3
	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	expected := 3
	assert.Equal(t, expected, findJudge(n, trust), "Judge id should be 3")
}

func TestThreeNodes(t *testing.T) {
	// input: n = 3, trust = [[1,3],[2,3],[3,1]]
	// output: -1
	n := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	expected := -1
	assert.Equal(t, expected, findJudge(n, trust), "Judge id should be -1")
}
