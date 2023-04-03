package main

import (
	"fmt"
	"math"
)

func getMaxValue(str string, k int) int {
	dp := make([][]int, len(str)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, i+1)
	}
	// 初始化
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	for i := 1; i < len(dp); i++ {
		dp[i][1] = count(str[:i])
	}
	for i := 1; i < len(dp); i++ {
		for j := 2; j <= k; j++ {
			for m := i; m >= j; m-- {
				dp[i][j] = min(dp[i][j], max(dp[m-1][j-1], count(str[m-1:i])))
			}
		}
	}
	return dp[len(str)][k]
}

// 计算str的权值
func count(str string) int {
	hashMap := make([]bool, 26)
	t := 0
	for i := 0; i < len(str); i++ {
		if !hashMap[str[i]-'a'] {
			t++
			hashMap[str[i]-'a'] = true
		}
	}
	return len(str) * t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMaxValue("ababbbb", 2))
}
