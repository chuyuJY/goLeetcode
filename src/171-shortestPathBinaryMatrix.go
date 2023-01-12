package main

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	res := 1
	queue := [][]int{[]int{0, 0}}
	nextQueue := [][]int{}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode[0] == len(grid)-1 && curNode[1] == len(grid[0])-1 {
			return res
		}
		for _, dir := range dirs {
			r, c := curNode[0]+dir[0], curNode[1]+dir[1]
			if 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0]) && grid[r][c] == 0 {
				nextQueue = append(nextQueue, []int{r, c})
				grid[r][c] = 1
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
