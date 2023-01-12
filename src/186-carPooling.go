package main

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	line := [][2]int{}
	for _, trip := range trips {
		line = append(line, [2]int{trip[1], trip[0]}, [2]int{trip[2], -trip[0]})
	}
	sort.Slice(line, func(i, j int) bool {
		return line[i][0] < line[j][0]
	})
	cnt := 0
	for _, point := range line {
		cnt += point[1]
		if cnt > capacity {
			return false
		}
	}
	return true
}
