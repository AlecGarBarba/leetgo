package inverttree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNulTree(t *testing.T) {
	var root *TreeNode
	invertedRoot := invertTree(root)

	assert.Nil(t, invertedRoot)
}

func TestOnlyRoot(t *testing.T) {
	root := &TreeNode{Val: 1}
	invertedRoot := invertTree(root)

	assert.NotNil(t, invertedRoot, "The root should not be nil")
	assert.Equal(t, 1, invertedRoot.Val, "The root value should be 1")
	assert.Nil(t, invertedRoot.Left, "The left child of the root should be nil")
	assert.Nil(t, invertedRoot.Right, "The right child of the root should be nil")
}

func TestOneLevelTree(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	invertedRoot := invertTree(root)

	assert.Equal(t, 1, invertedRoot.Val, "The root value should be 1")
	assert.Equal(t, 3, invertedRoot.Left.Val, "The left child value should be 3")
	assert.Equal(t, 2, invertedRoot.Right.Val, "The right child value should be 2")
}

func TestOneLevelImbalancedTree(t *testing.T) {
	/*
		Tree stucture:
			1
		   /
		  2
	*/
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	invertedRoot := invertTree(root)
	/*
		Tree structure after inversion:
			1
			 \
			  2
	*/
	assert.Equal(t, 1, invertedRoot.Val, "The root value should be 1")
	assert.Nil(t, root.Left, "The left child of the root should be nil")
	assert.Equal(t, 2, root.Right.Val, "The right child value should be 2")
}

func TestTwoLevelBalancedTree(t *testing.T) {
	/*
		Tree structure:
			1
		   / \
		  2   3
		 / \ / \
		4  5 6  7
	*/
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}}
	invertedRoot := invertTree(root)
	/**
	Tree structure after inversion:
		1
	   / \
	  3   2
	 / \ / \
	7  6 5  4
	*/
	assert.Equal(t, 1, invertedRoot.Val, "The root value should be 1")
	assert.Equal(t, 3, invertedRoot.Left.Val, "The left child value should be 3")
	assert.Equal(t, 2, invertedRoot.Right.Val, "The right child value should be 2")
	assert.Equal(t, 7, invertedRoot.Left.Left.Val, "The left child of the left child value should be 7")
	assert.Equal(t, 6, invertedRoot.Left.Right.Val, "The right child of the left child value should be 6")
	assert.Equal(t, 5, invertedRoot.Right.Left.Val, "The left child of the right child value should be 5")
	assert.Equal(t, 4, invertedRoot.Right.Right.Val, "The right child of the right child value should be 4")
}

func TestTwoLevelImbalancedTree(t *testing.T) {
	/*
		Tree
			1
		   /
		  2
		 /
		3
	*/
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	invertedRoot := invertTree(root)
	/*
		Tree after inversion
			1
			 \
			  2
			   \
			   	3
	*/
	assert.Equal(t, 1, invertedRoot.Val, "The root value should be 1")
	assert.Nil(t, invertedRoot.Left, "The left child of the root should be nil")
	assert.Equal(t, 2, invertedRoot.Right.Val, "invertedRoot")
	assert.Nil(t, invertedRoot.Right.Left, "The left child of the right child should be nil")
	assert.Equal(t, 3, invertedRoot.Right.Right.Val, "The right child value should be 3")

}
