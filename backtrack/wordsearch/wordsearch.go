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
		for row := 0; row < len(board[0]); row++ {
			if word[0] == board[col][row] {
				return dfs(board, word, col, row, 0)
			}
		}

	}

	return false
}

func dfs(board [][]byte, target string, currentColumn int, currentRow int, index int) bool {
	n := len(board)
	m := len(board[0])

	if index == len(target) {
		return true
	}

	// out of bounds, out of luck

	if currentColumn < 0 || currentColumn >= n {
		return false
	}

	if currentRow < 0 || currentRow >= m {
		return false
	}

	// if they are not the same, bad luck
	if board[currentColumn][currentRow] != target[index] {
		return false
	}

	return (dfs(board, target, currentColumn+1, currentRow, index+1) || dfs(board, target, currentColumn-1, currentRow, index+1) || dfs(board, target, currentColumn, currentRow+1, index+1) || dfs(board, target, currentColumn, currentRow-1, index+1))

}
