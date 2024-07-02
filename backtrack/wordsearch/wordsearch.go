package wordsearch

/*
*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) > len(board)*len(board[0]) {
		return false
	}

	for col := 0; col < len(board); col++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[col][j] {
				return dfs(board, word, col, j, 0)
			}
		}

	}

	return false
}

func dfs(board [][]byte, target string, i int, j int, index int) bool {
	n := len(board)
	m := len(board[0])

	if index == len(target) {
		return true
	}

	// out of bounds, out of luck

	if i < 0 || i >= n {
		return false
	}

	if j < 0 || j >= m {
		return false
	}

	// if they are not the same, bad luck
	if board[i][j] != target[index] {
		return false
	}

	return (dfs(board, target, i+1, j, index+1) || dfs(board, target, i-1, j, index+1) || dfs(board, target, i, j+1, index+1) || dfs(board, target, i, j-1, index+1))

}
