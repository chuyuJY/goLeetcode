package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscanln(sc, &n, &m, &k)
	bombs := [][]int{}
	for i := 0; i < k; i++ {
		var p, q int
		fmt.Fscanln(sc, &p, &q)
		bombs = append(bombs, []int{p - 1, q - 1})
	}
	var x1, y1, x2, y2 int
	fmt.Fscanln(sc, &x1, &y1, &x2, &y2)
	fmt.Println(test4(n, m, bombs, x1-1, y1-1, x2-1, y2-1))
}

func test4(n, m int, bombs [][]int, x1, y1, x2, y2 int) int {
	grid := make([][]int, n)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, m)
	}
	visitedMap := make([][]bool, len(grid))
	for i := 0; i < len(visitedMap); i++ {
		visitedMap[i] = make([]bool, len(grid[0]))
	}
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	// 多源 bfs 求炸弹到每个格子的最短距离
	queue := [][]int{}
	for _, bomb := range bombs {
		queue = append(queue, bomb)
		visitedMap[bomb[0]][bomb[1]] = true
	}
	nextQueue := [][]int{}
	dis := 0
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		grid[curNode[0]][curNode[1]] = dis
		for _, dir := range dirs {
			nextI, nextJ := curNode[0]+dir[0], curNode[1]+dir[1]
			if 0 <= nextI && nextI < len(grid) && 0 <= nextJ && nextJ < len(grid[0]) && !visitedMap[nextI][nextJ] {
				nextQueue = append(nextQueue, []int{nextI, nextJ})
				visitedMap[nextI][nextJ] = true
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = [][]int{}
			dis++
		}
	}

	// dfs 搜索该 limit 是否有效
	var dfs func(i, j int, limit int) bool
	dfs = func(i, j int, limit int) bool {
		// 剪枝
		if grid[i][j] < limit {
			return false
		}
		if i == x2 && j == y2 {
			return true
		}
		flag := false // 标识在此 limit 下, 是否可以到达
		visitedMap[i][j] = true
		for _, dir := range dirs {
			nextI, nextJ := i+dir[0], j+dir[1]
			if 0 <= nextI && nextI < len(grid) && 0 <= nextJ && nextJ < len(grid[0]) && !visitedMap[nextI][nextJ] {
				if dfs(nextI, nextJ, limit) {
					flag = true
					break
				}
			}
		}
		visitedMap[i][j] = false
		return flag
	}
	for i := 0; i < len(visitedMap); i++ {
		for j := 0; j < len(visitedMap[0]); j++ {
			visitedMap[i][j] = false
		}
	}
	//res := sort.Search(n+m-1, func(i int) bool {
	//	return !dfs(x1, y1, i)
	//})
	left, right := 0, n+m-1
	for left < right {
		mid := left + (right-left)/2
		if dfs(x1, y1, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}
