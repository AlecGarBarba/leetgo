package pathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTree(t *testing.T) {

	assert.False(t, hasPathSum(nil, 0))

}

func TestSingleNodeTreeFalse(t *testing.T) {
	root := &TreeNode{Val: 1}

	assert.False(t, hasPathSum(root, 0))

}

func TestSingleNodeTreeTrue(t *testing.T) {
	root := &TreeNode{Val: 1}

	assert.True(t, hasPathSum(root, 1))

}

func TestTwoLevelTreeFalse(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}

	assert.False(t, hasPathSum(root, 1))

}

func TestTwoLevelTreeTrue(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}

	assert.True(t, hasPathSum(root, 3), "Sum should equal 3")

}

func TestThreeLevelTreeFalse(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   3
	* /
	* 4
	 */
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}}

	assert.False(t, hasPathSum(root, 18))

}

func TestThreeLevelTreeTrue(t *testing.T) {
	/*
	* Tree Structure:
	*     1
	*    / \
	*   2   3
	*  /
	* 4
	 */

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	assert.True(t, hasPathSum(root, 7), "Sum should equal 7")

}
