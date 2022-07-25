package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	// 按照x坐标升序排列
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cover := points[0][1]
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= cover {
			// 更新最后边界
			if points[i][1] <= cover {
				cover = points[i][1]
			}
		} else {
			res++
			cover = points[i][1]
		}
	}
	return res
}

func main() {
	points := [][]int{{2, 5}, {1, 4}}
	fmt.Println(findMinArrowShots(points))
}
