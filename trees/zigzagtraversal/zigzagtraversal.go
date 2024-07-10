package zigzagtraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*

  - Problem statement:
    Given the root of a binary tree,
    return the zigzag level order traversal of its nodes' values.
    (i.e., from left to right, then right to left for the next level and alternate between).

    Example 1:
    Input: root = [3,9,20,null,null,15,7]
    Output: [[3],[20,9],[15,7]]
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		zig := len(res)%2 != 0

		for i := 0; i < levelSize; i++ {
			curr := queue[0]

			queue = queue[1:] // dequeue

			// do something with curr
			level = append(level, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		if zig && len(level) > 1 {
			level = reverseSlice(level)
		}

		res = append(res, level)

	}

	return res
}

func reverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return s
}
