package mindepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilTree(t *testing.T) {
	assert.Equal(t, 0, minDepth(nil), "Nil tree should return 0")
}

func TestOneNode(t *testing.T) {
	assert.Equal(t, 1, minDepth(&TreeNode{Val: 1}), "One node should return 1")
}

func TestTwoLevelBalancedTree(t *testing.T) {
	/* Tree Structure:
	 *     1
	 *    / \
	 *   2   3
	 */
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	assert.Equal(t, 2, minDepth(root), "Two level balanced tree should return 2")
}

func TestTwoLevelImbalancedTree(t *testing.T) {
	/* Tree
	 *     1
	 *    /
	 *   2
	 */
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	assert.Equal(t, 2, minDepth(root), "Two level imbalanced tree should return 2")
}

func TestThreeLevelBalancedTree(t *testing.T) {
	/*    Tree
	 *     1
	 *    / \
	 *   2   3
	 *  / \ / \
	 * 4  5 6  7
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	assert.Equal(t, 3, minDepth(root), "Three level balanced tree should return 3")
}

func TestThreeLevelImbalancedTree(t *testing.T) {
	/* Tree
	 *     1
	 *    /
	 *   2
	 *  /
	 * 3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	assert.Equal(t, 3, minDepth(root), "Three level imbalanced tree should return 3")
}
