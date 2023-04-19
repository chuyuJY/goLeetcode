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
	for i := 0; i < n-1; i++ {
		var s, t int
		fmt.Fscanln(sc, &s, &t)
		edges = append(edges, []int{s - 1, t - 1})
	}
	test4(n, edges)
}

func test4(n int, edges [][]int) {
	nodes := make([]int, n)
	graph := make([][]int, n)
	fathers := make([]int, n)
	inDegrees := make([]int, n)
	for _, edge := range edges {
		inDegrees[edge[1]]++
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		fathers[edge[0]] = edge[1]
	}
	queue := []int{}
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		nodes[cur] = 1
		for _, i := range graph[cur] {
			nodes[cur] += nodes[i]
		}
		inDegrees[fathers[cur]]--
		if inDegrees[fathers[cur]] == 0 {
			queue = append(queue, fathers[cur])
		}
	}
	res, count := n, -1
	for _, node := range nodes {
		cur := abs(n-node, node)
		if res > cur {
			res = cur
			count = 1
		} else if res == cur {
			count++
		}
	}
	fmt.Println(res, count)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
