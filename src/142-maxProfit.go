package main

func maxProfit(k int, prices []int) int {
	dp := make([][][2]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1)
	}
	for j := 0; j < k+1; j++ {
		dp[0][j][0] = -prices[0]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]+prices[i-1])
		}
	}
	return dp[len(dp)-1][k][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}