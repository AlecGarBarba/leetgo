package zigzagtraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilNode(t *testing.T) {
	var root *TreeNode
	assert.Equal(t, [][]int{}, zigzagLevelOrder(root))

}

func TestSingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	assert.Equal(t, [][]int{{1}}, zigzagLevelOrder(root))
}

func TestTwoLevelBalancedTree(t *testing.T) {
	/**
	 *    1
	 *	 / \
	 *  2   3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	assert.Equal(t, [][]int{{1}, {3, 2}}, zigzagLevelOrder(root))
}

func TestTwoLevelImbalancedTree(t *testing.T) {
	/**
	 *    1
	 *	 /
	 *  2
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	assert.Equal(t, [][]int{{1}, {2}}, zigzagLevelOrder(root))
}

func TestThreeLevelBalancedTree(t *testing.T) {
	/*
	 *    1
	 *	 / \
	 *  2   3
	 * / \ / \
	 * 4 5 6 7
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}, zigzagLevelOrder(root))
}

func TestThreeLevelImbalancedTree(t *testing.T) {
	/*
	 *    1
	 *	 /
	 *  2
	 * /
	 * 3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	assert.Equal(t, [][]int{{1}, {2}, {3}}, zigzagLevelOrder(root))
}
