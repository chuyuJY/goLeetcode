package main

import "sort"

func merge(intervals [][]int) [][]int {
	// 升序排列，按照左区间
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
