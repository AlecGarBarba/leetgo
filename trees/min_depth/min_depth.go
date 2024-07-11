package mindepth

/**
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	/*
		Base Case 1 - Empty Tree: If the root is nil, it means the tree is empty, so the depth is 0.
	*/
	if root == nil {
		return 0
	}

	/*
		Base Case 2 - Leaf Node: If both the left and right child nodes of the root are nil, it means the root is a leaf node, so the depth is 1.
	*/
	if root.Left == nil && root.Right == nil {
		return 1
	}

	/*
		Recursive Case 1 - Only Right Subtree: If the left child node is nil, it recursively calculates the minimum depth of the right subtree and adds 1 to the result. The +1 is to account for the current node.
	*/
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	/*
		Recursive Case 2 - Only Left Subtree: If the right child node is nil, it recursively calculates the minimum depth of the left subtree and adds 1 to the result. The +1 is to account for the current node.
	*/
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	/*
		Recursive Case 3 - Both Subtrees: If both child nodes are not nil, it recursively calculates the minimum depths of both the left and right subtrees, and returns the smaller one plus 1. The +1 is to account for the current node.
	*/
	left := minDepth(root.Left) + 1
	right := minDepth(root.Right) + 1

	if left < right {
		return left
	}

	return right
}
