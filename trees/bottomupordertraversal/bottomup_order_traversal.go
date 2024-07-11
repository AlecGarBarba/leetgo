package bottomupordertraversal

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

func levelOrderBottom(root *TreeNode) [][]int {

	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	stack := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:] // dequeue

			level = append(level, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		stack = append(stack, level)

	}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		res = append(res, last)
		// pop
		stack = stack[:len(stack)-1]
	}

	return res
}
