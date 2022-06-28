package main

import "math"

func updateMatrix(mat [][]int) [][]int {
	dists := make([][]int, len(mat))
	queue := [][]int{}
	for i := 0; i < len(dists); i++ {
		dists[i] = make([]int, len(mat[i]))
		for j := 0; j < len(dists[i]); j++ {
			if mat[i][j] == 0 {
				dists[i][j] = 0
				queue = append(queue, []int{i, j})
			} else {
				dists[i][j] = math.MaxInt
			}
		}
	}
	bfsMatrix(mat, dists, queue)
	return dists
}

func bfsMatrix(mat [][]int, dists [][]int, queue [][]int) {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curDist := dists[cur[0]][cur[1]]
		for _, dir := range dirs {
			r := cur[0] + dir[0]
			c := cur[1] + dir[1]
			if r >= 0 && r < len(mat) && c >= 0 && c < len(mat[r]) {
				if dists[r][c] > curDist+1 {
					dists[r][c] = curDist + 1
					queue = append(queue, []int{r, c})
				}
			}
		}
	}
}
