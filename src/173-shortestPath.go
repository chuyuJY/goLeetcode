package main

func shortestPath(grid [][]int, k int) int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]int, len(grid)) // 标记到达当前位置时的剩余k值
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			visited[i][j] = -1
		}
	}
	// BFS
	res := 0
	visited[0][0] = k
	queue := [][]int{{0, 0}}
	nextQueue := [][]int{}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode[0] == len(grid)-1 && curNode[1] == len(grid[0])-1 {
			return res
		}
		for _, dir := range dirs {
			r := curNode[0] + dir[0]
			c := curNode[1] + dir[1]
			if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0]) {
				if grid[r][c] == 0 {
					if visited[r][c] < visited[curNode[0]][curNode[1]] {
						nextQueue = append(nextQueue, []int{r, c})
						visited[r][c] = visited[curNode[0]][curNode[1]]
					}
				} else {
					if visited[r][c] < visited[curNode[0]][curNode[1]]-1 {
						nextQueue = append(nextQueue, []int{r, c})
						visited[r][c] = visited[curNode[0]][curNode[1]] - 1
					}
				}
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = [][]int{}
			res++
		}
	}
	return -1
}
