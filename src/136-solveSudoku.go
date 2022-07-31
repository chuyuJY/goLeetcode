package main

func solveSudoku(board [][]byte) {
	var backTracking func() bool
	backTracking = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] != '.' {
					continue
				}
				for num := '1'; num <= '9'; num++ {
					if isValid(byte(num), i, j, board) {
						board[i][j] = byte(num)
						if backTracking() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backTracking()
}

func isValid(num byte, row, col int, board [][]byte) bool {
	// 1. 检查同行和同列
	for i := 0; i < len(board); i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}
	// 2. 检查九宫格
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}
