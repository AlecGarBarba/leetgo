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
			if word[0] == board[col][row] && dfs(board, word, col, row, 0) {
				return true
			}
		}

	}

	return false
}

func dfs(board [][]byte, target string, currentColumn int, currentRow int, index int) bool {
	columns := len(board)
	rows := len(board[0])

	if index == len(target) {
		return true
	}

	if arePointersOutOfBounds(currentRow, currentColumn, columns, rows) || !isCurrentPositionMatch(board[currentColumn][currentRow], target[index]) {
		return false
	}

	return (dfs(board, target, currentColumn+1, currentRow, index+1) || dfs(board, target, currentColumn-1, currentRow, index+1) || dfs(board, target, currentColumn, currentRow+1, index+1) || dfs(board, target, currentColumn, currentRow-1, index+1))

}

func arePointersOutOfBounds(currRow int, currCol int, totalCols int, totalRows int) bool {
	if currCol < 0 || currCol >= totalCols {
		return true
	}

	if currRow < 0 || currRow >= totalRows {
		return true
	}

	return false
}

func isCurrentPositionMatch(boardValue byte, targetValue byte) bool {
	return boardValue == targetValue
}
