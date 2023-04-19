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
	var edgeNum int
	fmt.Fscanln(sc, &edgeNum)
	edges := [][]int{}
	for i := 0; i < edgeNum; i++ {
		var x, y int
		fmt.Fscanln(sc, &x, &y)
		edges = append(edges, []int{x, y})
	}
	var blockNum int
	fmt.Fscanln(sc, &blockNum)
	blocks := []int{}
	for i := 0; i < blockNum; i++ {
		var x int
		fmt.Fscanln(sc, &x)
		blocks = append(blocks, x)
	}
	test5(n, edges, blocks)
}

func test5(n int, edges [][]int, blocks []int) {
	// 1. 建图
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	// 2. block
	blockMap := map[int]bool{}
	for _, block := range blocks {
		blockMap[block] = true
	}
	if blockMap[0] {
		fmt.Println("NULL")
		return
	}
	// 3. BFS 先得到最短长度
	queue := []int{0}
	nextQueue := []int{}
	flag := false
	length := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 到达叶子节点
		if len(graph[cur]) == 0 {
			flag = true
			break
		}
		for _, next := range graph[cur] {
			if _, ok := blockMap[next]; !ok {
				nextQueue = append(nextQueue, next)
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = []int{}
			length++
		}
	}
	if !flag {
		fmt.Println("NULL")
		return
	}
	// 4. dfs 寻找路径
	path := []int{}
	var dfs func(curNode int) bool
	dfs = func(curNode int) bool {
		path = append(path, curNode)
		if len(path) == length && len(graph[curNode]) == 0 {
			return true
		}
		// 剪枝
		if len(path) == length {
			path = path[:len(path)-1]
			return false
		}
		for _, next := range graph[curNode] {
			if _, ok := blockMap[next]; !ok {
				if dfs(next) {
					return true
				}
			}
		}
		path = path[:len(path)-1]
		return false
	}
	dfs(0)
	for i := 0; i < len(path); i++ {
		if i < len(path)-1 {
			fmt.Printf("%v->", path[i])
		} else {
			fmt.Printf("%v", path[i])
		}
	}
}
