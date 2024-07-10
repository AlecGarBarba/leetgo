package sametree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	right := &TreeNode{Val: 3}
	left := &TreeNode{Val: 3}
	assert.True(t, IsSameTree(left, right))

}

func TestRootDifferent(t *testing.T) {
	right := &TreeNode{Val: 2}
	left := &TreeNode{Val: 3}
	assert.False(t, IsSameTree(left, right))

}

func TestOneIsNil(t *testing.T) {
	right := &TreeNode{Val: 3}
	assert.False(t, IsSameTree(nil, right))

}

func TestSameStructureTreeOneLevel(t *testing.T) {
	// input: [0,-5], [0,-8]
	/* Trees Structure:
	 *    0        0
	 *   /        /
	 * -5      -8
	 */
	left := &TreeNode{Val: 0, Left: &TreeNode{Val: -5}}
	right := &TreeNode{Val: 0, Left: &TreeNode{Val: -8}}
	assert.False(t, IsSameTree(left, right))
}

func TestNullTrees(t *testing.T) {
	assert.True(t, IsSameTree(nil, nil))
}

func TestThreeLevelTrees(t *testing.T) {
	left := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	right := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	assert.True(t, IsSameTree(left, right))

}

func TestImbalancedTrees(t *testing.T) {
	left := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	right := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}

	assert.True(t, IsSameTree(left, right))
}

func TestImbalancedDifferentTrees(t *testing.T) {
	left := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	right := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	assert.False(t, IsSameTree(left, right))
}
