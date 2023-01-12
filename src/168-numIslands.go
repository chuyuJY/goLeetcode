package main

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func bfs(grid [][]byte, i, j int) {
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][2]int{{i, j}}
	grid[i][j] = '0'
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			r, c := curNode[0]+dir[0], curNode[1]+dir[1]
			if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0]) && grid[r][c] == '1' {
				grid[r][c] = '0'
				queue = append(queue, [2]int{r, c})
			}
		}
	}
}
