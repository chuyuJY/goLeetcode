package main

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([]int, len(envelopes)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			if envelopes[j-1][1] < envelopes[i-1][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
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

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	maxEnvelopes(envelopes)
}
