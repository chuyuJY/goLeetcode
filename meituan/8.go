package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxValue([]int{1, 2, 3, 4, 5, 6, 7}, 1))
}

func maxValue(nums []int, k int) int {
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, k+1)
	}
	// 初始化
	for j := 0; j <= k; j++ {
		dp[1][j] = nums[0]
	}
	for i := 2; i < len(dp); i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-2][j]+nums[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], nums[i-1]+max(dp[i-2][j], dp[i-1][j-1]))
			}
		}
	}
	return dp[len(nums)][k]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
