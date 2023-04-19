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
	edges := [][]int{}
	for i := 1; i < n; i++ {
		var pi int
		fmt.Fscan(sc, &pi)
		edges = append(edges, []int{i + 1, pi})
	}
	fmt.Fscanln(sc)
	var x, y int
	fmt.Fscanln(sc, &x, &y)
	test14(n, edges, x, y)
}

func test14(n int, edges [][]int, x, y int) {
	// 建图
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// 分别从 x, y 向两边延伸, 分别得到最长的, 相加 + 1
	visitedMap := map[int]bool{}
	var dfs func(node int) int
	dfs = func(node int) int {
		if len(graph[node]) == 0 {
			return 1
		}
		res := 0
		for _, next := range graph[node] {
			if !visitedMap[next] {
				visitedMap[next] = true
				res = max(res, 1+dfs(next))
				visitedMap[next] = false
			}
		}
		return res
	}
	visitedMap[x] = true
	visitedMap[y] = true
	resX, resY := dfs(x), dfs(y)
	fmt.Println(resX + resY + 1)
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
