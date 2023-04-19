package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	str, _ := sc.ReadString('\n')
	colors := strings.TrimRight(str, "\r\n")
	colors = "#" + colors
	edges := [][]int{}
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(sc, &u, &v)
		edges = append(edges, []int{u, v})
	}
	test5(n, colors, edges)
}

func test5(n int, colors string, edges [][]int) {
	// 1. 计算红、蓝的总数
	sumRed, sumBlue := 0, 0
	for i := 1; i < len(colors); i++ {
		if colors[i] == 'R' {
			sumRed++
		} else {
			sumBlue++
		}
	}
	// 2. 建图
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	// 3. dfs
	res := 0
	var dfs func(curNode int) (int, int)
	dfs = func(curNode int) (int, int) {
		curRed, curBlue := 0, 0
		if len(graph[curNode]) > 0 {
			tempRed, tempBlue := 0, 0
			for _, next := range graph[curNode] {
				tempRed, tempBlue = dfs(next)
				curRed += tempRed
				curBlue += tempBlue
			}
		}
		if colors[curNode] == 'R' {
			curRed++
		} else {
			curBlue++
		}
		if sumRed-curRed > sumBlue-curBlue && curRed > curBlue {
			res++
		}
		return curRed, curBlue
	}
	dfs(1)
	fmt.Println(res)
}
