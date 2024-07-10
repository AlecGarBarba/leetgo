package maxdepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilNode(t *testing.T) {
	assert.Equal(t, 0, maxDepth(nil), "Should be 0 for nil node")

}

func TestOneNode(t *testing.T) {
	node := &TreeNode{Val: 1}
	assert.Equal(t, 1, maxDepth(node), "Should be 1 for one node")
	node2 := &TreeNode{}
	assert.Equal(t, 1, maxDepth(node2))

}

func TestTwoLevelBalancedTree(t *testing.T) {
	node := &TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 3}
	assert.Equal(t, 2, maxDepth(node), "Should be 2 for two level balanced tree")
}

func TestThreevelImbalancedTree(t *testing.T) {
	node := &TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Left.Left = &TreeNode{Val: 3}
	assert.Equal(t, 3, maxDepth(node), "Should be 3 for two level imbalanced tree")
}
