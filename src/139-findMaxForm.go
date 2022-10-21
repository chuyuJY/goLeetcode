package main

func findMaxForm(strs []string, m int, n int) int {
	// 数据预处理
	costs := make([][2]int, len(strs))
	for i := 0; i < len(costs); i++ {
		costs[i] = cnt(strs[i])
	}
	// 动态规划
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for _, cost := range costs {
		for i := m; i >= cost[0]; i-- {
			for j := n; j >= cost[1]; j-- {
				dp[i][j] = max(dp[i][j], dp[i-cost[0]][j-cost[1]]+1)
			}
		}
	}
	return dp[m][n]
}

func cnt(s string) [2]int {
	x, y := 0, 0
	for _, char := range s {
		switch char {
		case '0':
			x++
		case '1':
			y++
		}
	}
	return [2]int{x, y}
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
