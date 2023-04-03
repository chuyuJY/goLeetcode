package main

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][3]int, len(startTime))
	for i := 0; i < len(jobs); i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	dp := make([]int, len(jobs)+1)

	res := 0
	for i := 1; i < len(dp); i++ {
		// 二分查找前一个
		preIndex := sort.Search(len(jobs), func(j int) bool {
			return jobs[j][1] <= jobs[i-1][0]
		})
		if preIndex > -1 {
			dp[i] = max(dp[i-1], dp[preIndex]+jobs[i-1][2])
		} else {
			dp[i] = max(dp[i-1], profit[i-1])
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
