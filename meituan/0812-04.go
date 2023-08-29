package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	var str string
	fmt.Fscanln(sc, &str)
	test19(str)
}

func test19(s string) {
	res := len(s)
	for i := 1; i < len(s); i++ {
		if len(s)%i == 0 {
			res = min(res, calUnion(buildGrid(s, i, len(s)/i)))
		}
	}
	fmt.Println(res)
}

func buildGrid(s string, row, col int) [][]byte {
	grid := make([][]byte, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]byte, col)
	}
	for i := 0; i < len(s); i++ {
		grid[i/col][i%col] = s[i]
	}
	return grid
}

func calUnion(grid [][]byte) int {
	res := 0
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	var dfs func(i, j int, tag byte)
	dfs = func(i, j int, tag byte) {
		visited[i][j] = true
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == tag && !visited[x][y] {
				dfs(x, y, tag)
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				dfs(i, j, grid[i][j])
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
