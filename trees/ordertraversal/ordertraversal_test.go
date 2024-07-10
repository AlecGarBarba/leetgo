package ordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilNodeTraversal(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil))

}

func TestSingleNodeTraversal(t *testing.T) {
	node := &TreeNode{Val: 1}

	assert.Equal(t, [][]int{{1}}, levelOrder(node))
}

func TestTwoLevelTraversal(t *testing.T) {
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	assert.Equal(t, [][]int{{1}, {2, 3}}, levelOrder(node))

}
