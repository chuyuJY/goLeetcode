package main

import "fmt"

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	res := 0

	dp := make([][2]int, len(nums1))
	var dfs func(curVal, cur int) int
	dfs = func(curVal, cur int) int {
		if cur == len(nums1)-1 {
			return 1
		}
		curRes := 1
		nextMin, nextMax := min(nums1[cur+1], nums2[cur+1]), max(nums1[cur+1], nums2[cur+1])
		// 重开
		if nextMin <= curVal {
			if dp[cur+1][0] == 0 {
				dp[cur+1][0] = dfs(nextMin, cur+1)
			}
			res = max(res, dp[cur+1][0])
		}
		// 当前
		if curVal <= nextMin {
			if dp[cur+1][0] == 0 {
				dp[cur+1][0] = dfs(nextMin, cur+1)
			}
			curRes += dp[cur+1][0]
		} else if curVal <= nextMax {
			if dp[cur+1][1] == 0 {
				dp[cur+1][1] = dfs(nextMax, cur+1)
			}
			curRes += dp[cur+1][1]
		}
		return curRes
	}
	res = max(res, dfs(min(nums1[0], nums2[0]), 0))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxNonDecreasingLength([]int{15, 108, 15, 41, 15, 12, 15, 51, 15, 78, 15, 104, 15, 19, 15, 74, 15, 48, 15, 34, 15, 64, 15, 51, 15, 22, 15, 9, 15, 98, 15}, []int{83, 15, 60, 15, 106, 15, 87, 15, 37, 15, 59, 15, 8, 15, 83, 15, 74, 15, 105, 15, 47, 15, 3, 15, 77, 15, 34, 15, 67, 15, 90}))
}
