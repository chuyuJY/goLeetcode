package main

import "sort"

func longestStrChain(words []string) int {
	dp := make([]int, len(words))
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if isValid(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func isValid(a, b string) bool {
	if len(b)-len(a) != 1 {
		return false
	}
	indexA, indexB := 0, 0
	flag := false
	for indexA < len(a) {
		if a[indexA] == b[indexB] {
			indexA++
			indexB++
			continue
		}
		if flag {
			return false
		}
		flag = true
		indexB++
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
