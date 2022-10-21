package main

func longestIncreasingPath(matrix [][]int) int {
	res := 0
	// 记忆矩阵
	length := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res = max(res, dfs(i, j, matrix, length))
		}
	}
	return res
}

func dfs(i, j int, matrix [][]int, length [][]int) int {
	if length[i][j] > 0 {
		return length[i][j]
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[newI]) && matrix[i][j] < matrix[newI][newJ] {
			length[i][j] = max(length[i][j], 1+dfs(newI, newJ, matrix, length))
		}
	}
	return length[i][j]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
