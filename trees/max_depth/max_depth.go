package maxdepth

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

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	count := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		count++

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:] // dequeue

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}

		}
	}

	return count
}
