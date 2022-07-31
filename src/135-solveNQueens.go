package main

func solveNQueens(n int) [][]string {
	res := [][]string{}
	// 1. 初始化棋盘
	chessBoard := make([][]byte, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = '.'
		}
	}
	// 2. 回溯
	var backTracking func(row int)
	backTracking = func(row int) {
		if row == n {
			path := make([]string, n)
			for i := 0; i < n; i++ {
				path[i] = string(chessBoard[i])
			}
			res = append(res, path)
			return
		}

		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessBoard) {
				chessBoard[row][i] = 'Q'
				backTracking(row + 1)
				chessBoard[row][i] = '.'
			}
		}
	}
	backTracking(0)
	return res
}

func isValid(n, row, col int, chessBoard [][]byte) bool {
	// 1. 检查列
	for i := 0; i < n; i++ {
		if i != row && chessBoard[i][col] == 'Q' {
			return false
		}
	}

	// 2. 检查斜角线
	// 检查左上角
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if chessBoard[r][c] == 'Q' {
			return false
		}
	}
	// 检查右上角
	for r, c := row-1, col+1; r >= 0 && c <= n-1; r, c = r-1, c+1 {
		if chessBoard[r][c] == 'Q' {
			return false
		}
	}
	return true
}
