package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	var dfs func(cur int, path []int)
	dfs = func(cur int, path []int) {
		if cur == len(graph)-1 {
			res = append(res, path)
			return
		}
		for _, node := range graph[cur] {
			path = append(path, node)
			dfs(node, append([]int{}, path...))
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{0})
	return res
}

func main() {
	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
