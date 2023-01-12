package main

import "math"

func updateMatrix(mat [][]int) [][]int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	dists := make([][]int, len(mat))
	queue := [][]int{}
	for i := 0; i < len(dists); i++ {
		dists[i] = make([]int, len(mat[i]))
		for j := 0; j < len(dists[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
				dists[i][j] = 0
			} else {
				dists[i][j] = math.MaxInt
			}
		}
	}
	// BFS
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		curDist := dists[curNode[0]][curNode[1]]
		for _, dir := range dirs {
			r := curNode[0] + dir[0]
			c := curNode[1] + dir[1]
			if 0 <= r && r < len(mat) && 0 <= c && c < len(mat[0]) {
				if dists[r][c] > curDist+1 {
					dists[r][c] = curDist + 1
					queue = append(queue, []int{r, c})
				}
			}
		}
	}
	return dists
}
