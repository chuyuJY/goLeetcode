package main

func minCost(costs [][]int) int {
	dp := make([][3]int, len(costs))
	// 初始化
	for j := 0; j < 3; j++ {
		dp[0][j] = costs[0][j]
	}
	for i := 1; i < len(costs); i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = costs[i][j] + min(dp[i][(j+1)%2], dp[i][(j+2)%2])
		}
	}
	return min(min(dp[len(costs)-1][0], dp[len(costs)-1][1]), dp[len(costs)-1][2])
}
