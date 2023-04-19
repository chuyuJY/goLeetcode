package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewReader(os.Stdin)
	str, _ := sc.ReadString('\n')
	colors := strings.TrimRight(str, "\r\n")

	edges := [][]int{}
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(sc, &u, &v)
		edges = append(edges, []int{u - 1, v - 1})
	}
	
	fmt.Println(test2(colors, edges))
}

func test2(colors string, edges [][]int) int {
	graph := map[int][]int{}
	inDegrees := make([]int, len(colors))
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		inDegrees[edge[1]]++
	}

	return min(updateColor(graph, inDegrees, colors, 0), updateColor(graph, inDegrees, colors, 1))
}

func updateColor(graph map[int][]int, inDegrees []int, colors string, startColor int) int {
	res := 0
	colorMap := []byte{'R', 'W'}
	tmpInDegrees := make([]int, len(inDegrees))
	copy(tmpInDegrees, inDegrees)
	queue := []int{}
	nextQueue := []int{}
	for i, inDegree := range tmpInDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
		}
	}
	curColor := startColor
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if colors[cur] != colorMap[curColor] {
			res++
		}
		for _, node := range graph[cur] {
			tmpInDegrees[node]--
			if tmpInDegrees[node] == 0 {
				nextQueue = append(nextQueue, node)
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = []int{}
			curColor ^= 1
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
