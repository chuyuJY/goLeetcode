package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 二维数组排序, 按照第一个位置
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 新建数组
	res := [][]int{}
	// 滑动窗口
	left, right := 0, 0
	for right < len(intervals) {
		new := []int{intervals[left][0], intervals[left][1]}
		// 滑动窗口右移
		for right < len(intervals) && intervals[right][0] <= new[1] {
			new[1] = max(new[1], intervals[right][1])
			right++
		}
		res = append(res, new)
		left = right
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
