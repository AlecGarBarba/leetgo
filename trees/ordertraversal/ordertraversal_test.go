package ordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilNodeTraversal(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrder(nil), "Should return empty array when root of the tree is nil")

}

func TestSingleNodeTraversal(t *testing.T) {
	node := &TreeNode{Val: 1}

	assert.Equal(t, [][]int{{1}}, levelOrder(node), "Should return single value array for a single node tree")
}

func TestTwoLevelTreeTraversal(t *testing.T) {
	/* Given the following tree structure
	 *     1
	 *    / \
	 *   2   3
	 */
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	assert.Equal(t, [][]int{{1}, {2, 3}}, levelOrder(node), "Should return two level traversal")
}

func TestTrheeLevelTreeTraversal(t *testing.T) {
	/* Given the following tree structure
	 *     1
	 *    / \
	 *   2   3
	 *  / \ / \
	 * 4  5 6  7
	 */
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}

	assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}, levelOrder(node), "Should return three level traversal")
}
