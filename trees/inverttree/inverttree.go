package inverttree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Invert a binary tree, given the root of the tree.
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// switch current level left and right
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	// invert left and right subtree
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
