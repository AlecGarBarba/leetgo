package issimmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullTrue(t *testing.T) {

	assert.True(t, isSymmetric(nil), "Should return true for nil tree")

}

func TestOneNodeTrue(t *testing.T) {

	root := &TreeNode{Val: 1}

	assert.True(t, isSymmetric(root), "Should return true for one node tree")

}

func TestSimmetricTwoLevelTree(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   2
	 */

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}

	assert.True(t, isSymmetric(root), "Should return true for two level simmetric tree")

}

func TestAssimetricTwoLevelTree(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	assert.False(t, isSymmetric(root), "Should return false for two level assimetric tree")
}

func TestAssimetricTwoLevelTreeWithNil(t *testing.T) {
	/*
	* Tree Structure:
	*     0
	*    /
	*   1
	 */
	root := &TreeNode{Val: 0, Left: &TreeNode{Val: 1}}
	assert.False(t, isSymmetric(root), "Should return false for two level assimetric tree with nil")

}

func TestSimmetricThreeLevelTree(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   2
	*  / \ / \
	* 3  4 4  3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	assert.True(t, isSymmetric(root), "Should return true for three level simmetric tree")
}

func TestAssimmetricThreeLevelTree(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   2
	*  / \ / \
	* 3  4 3  4
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}
	assert.False(t, isSymmetric(root), "Should return false for three level assimetric tree")
}

func TestAssimetricThreeLevelTreeWithNullValues(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   2
	*  / \ / \
	* 3  4 4  nil
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}}
	assert.False(t, isSymmetric(root), "Should return false for three level assimetric tree with null values")
}

func TestThreeLevelTreeTwoNil(t *testing.T) {
	// input is [1,2,2,null,3,3]

	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   2
	*    \   \
	*     3   3
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	assert.False(t, isSymmetric(root), "Should return false for three level assimetric tree with two nil values")
}
