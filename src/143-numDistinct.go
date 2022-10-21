package main

import "fmt"

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}
