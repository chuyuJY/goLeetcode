package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := intervals[0]
	for _, interval := range intervals {
		if cur[1] < interval[0] {
			res = append(res, cur)
			cur = interval
		} else {
			cur[0] = min(cur[0], interval[0])
			cur[1] = max(cur[1], interval[1])
		}
	}
	res = append(res, cur)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
