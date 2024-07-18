package pathsum

/**
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

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return dfsSum(root.Left, targetSum, root.Val) || dfsSum(root.Right, targetSum, root.Val)
}

func dfsSum(root *TreeNode, targetSum int, current int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		total := root.Val + current
		return targetSum == total
	}

	return dfsSum(root.Left, targetSum, current+root.Val) || dfsSum(root.Right, targetSum, current+root.Val)
}
