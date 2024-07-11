package bottomupordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTree(t *testing.T) {
	assert.Equal(t, [][]int{}, levelOrderBottom(nil))

}

func TestOnlyRootNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := [][]int{{1}}
	assert.Equal(t, expected, levelOrderBottom(root), "Should return the root node value")
}

func TestTwoLevelTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	expected := [][]int{{2, 3}, {1}}
	assert.Equal(t, expected, levelOrderBottom(root), "Should return the root node value")
}

func TestThreeLevelImbalancedTree(t *testing.T) {
	/** Tree Structure
	 *     1
	 *    / \
	 *   2   3
	 *  / \
	 * 4   5
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	expected := [][]int{{4}, {2, 3}, {1}}
	assert.Equal(t, expected, levelOrderBottom(root), "Should return the root node value")
}

func TestThreeLevelBalancedTree(t *testing.T) {
	/** Tree Structure
	 *     1
	 *    / \
	 *   2   3
	 *  / \ / \
	 * 4  5 6  7
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	expected := [][]int{{4, 5, 6, 7}, {2, 3}, {1}}
	assert.Equal(t, expected, levelOrderBottom(root), "Should return the root node value")
}

/*
* Leetcode Test case
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]
*/
func TestExample1(t *testing.T) {
	/* Tree Structure
	 *     3
	 *    / \
	 *   9  20
	 *     /  \
	 *    15   7
	 */
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := [][]int{{15, 7}, {9, 20}, {3}}
	assert.Equal(t, expected, levelOrderBottom(root), "Should return the root node value")
}
