package issimmetric

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	return areNodesSimmetric(root.Left, root.Right)
}

func areNodesSimmetric(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if (left != nil && right == nil) || (left == nil && right != nil) {
		return false
	}

	return left.Val == right.Val && areNodesSimmetric(left.Left, right.Right) && areNodesSimmetric(left.Right, right.Left)
}
