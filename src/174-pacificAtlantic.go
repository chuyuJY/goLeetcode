package main

func pacificAtlantic(heights [][]int) [][]int {
	res := [][]int{}
	res1, res2 := make([][]bool, len(heights)), make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		res1[i], res2[i] = make([]bool, len(heights[i])), make([]bool, len(heights[i]))
	}
	// 遍历太平洋海岸
	for i := 0; i < len(heights); i++ {
		bfs(res1, heights, i, 0)
	}
	for j := 0; j < len(heights[0]); j++ {
		bfs(res1, heights, 0, j)
	}
	// 遍历大西洋海岸
	for i := 0; i < len(heights); i++ {
		bfs(res2, heights, i, len(heights[0])-1)
	}
	for j := 0; j < len(heights[0]); j++ {
		bfs(res2, heights, len(heights)-1, j)
	}
	// 取交集
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if res1[i][j] && res2[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(res [][]bool, heights [][]int, i, j int) {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][]int{{i, j}}
	res[i][j] = true
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			r := curNode[0] + dir[0]
			c := curNode[1] + dir[1]
			if 0 <= r && r < len(heights) && 0 <= c && c < len(heights[0]) {
				if !res[r][c] && heights[r][c] >= heights[curNode[0]][curNode[1]] {
					queue = append(queue, []int{r, c})
					res[r][c] = true
				}
			}
		}
	}
}
