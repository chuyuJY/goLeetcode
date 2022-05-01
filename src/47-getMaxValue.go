package main

import "fmt"

func getMaxValue(nums []int, k int) int {
	cnt := len(nums)
	// 初始化dp
	dp := make([][]int, cnt+1)
	for i := 0; i <= cnt; i++ {
		if i == 0 {
			dp[i] = make([]int, 1)
		} else {
			dp[i] = make([]int, i)
		}
	}
	//
	for i := 1; i <= cnt; i++ {
		dp[i][0] += nums[i-1] + dp[i-1][0]
	}
	// 先列
	for j := 1; j <= k; j++ {
		// 行
		for i := 2; i <= cnt; i++ {
			// 数字比乘号多
			if i > j {
				for p := 1; p < i; p++ {
					// 乘号没用完的情况 是不可能为最优解的
					if p > j-1 {
						// dp[p][j-1]是子问题的最优解
						dp[i][j] = max(dp[i][j], dp[p][j-1]*(dp[i][0]-dp[p][0]))
					}
				}
			}
		}
	}
	return dp[cnt][k]
}

//func minMaxSum(nums []int, k int) int {
//
//}

//func max(x int, y int) int {
//	if x >= y {
//		return x
//	} else {
//		return y
//	}
//}

//func min(x int, y int) int {
//	if x <= y {
//		return x
//	} else {
//		return y
//	}
//}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(getMaxValue(arr, 2))
	//arr = []int{2, 2, 2, 8, 1, 8, 2, 1}
	//fmt.Println(minMaxSum(arr, 17))
}
