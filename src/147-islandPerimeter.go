package main

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 只统计陆地块
			if grid[i][j] == 1 {
				res += countEdges(grid, i, j)
			}
		}
	}
	return res
}

func countEdges(grid [][]int, i, j int) int {
	edges := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
			edges++
		}
	}
	return edges
}
