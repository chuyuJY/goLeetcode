package main

func getWhiteCounts(rects [][]int) int {
	graph := make([][]int, 100)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 100)
		for j := 0; j < 100; j++ {
			graph[i][j] = 1
		}
	}
	for _, rect := range rects {
		startX, startY := min(rect[0], rect[2]), min(rect[1], rect[3])
		endX, endY := max(rect[0], rect[2]), max(rect[1], rect[3])
		for i := startX; i < endX; i++ {
			for j := startY; j < endY; j++ {
				graph[i][j] = -graph[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == 1 {
				res++
			}
		}
	}
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
