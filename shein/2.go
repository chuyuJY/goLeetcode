package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)
	fmt.Println(test2(n, m, x, y))
}

func test2(n, m, x, y int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	if !isValid(0, 0, x, y) {
		return 0
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		if isValid(i, 0, x, y) {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j <= m; j++ {
		if isValid(0, j, x, y) {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if isValid(i, j, x, y) {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

func isValid(x1, y1 int, x, y int) bool {
	if (x1 == x && y1 == y) || (x1 != x && y1 != y && abs(x1-x)+abs(y1-y) == 3) {
		return false
	}
	return true
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
