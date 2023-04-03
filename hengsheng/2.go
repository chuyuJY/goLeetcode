package main

func getMax(M float64, N int, historyPrices []float64, k int) {
	dp := make([][][2]int, len(historyPrices)+1)
	count, money := 0, M
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k) // 0: 不持有, 1: 持有
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

