package ordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

example:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
*/

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:] // dequeue

			level = append(level, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}

		}

		result = append(result, level)
	}

	return result

}
