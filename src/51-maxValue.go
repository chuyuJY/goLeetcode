package main

func maxValue(grid [][]int) int {
	l, w := len(grid), len(grid[0])
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, w)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < l; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < w; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < l; i++ {
		for j := 1; j < w; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[l-1][w-1]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
