package main

import "fmt"

func isBipartite(graph [][]int) bool {
	// 标记节点染色情况, 0-未染色, -1-黑, 1-白
	colors := make([]int, len(graph))
	// 每轮染色一个子图
	for i := 0; i < len(colors); i++ {
		if colors[i] == 0 {
			// 对该点所在图染色失败
			if !setColor(graph, colors, i, 1) {
				return false
			}
		}
	}
	return true
}

// 广度优先搜索
func setColor(graph [][]int, colors []int, i, color int) bool {
	colors[i] = color
	queue := []int{i}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, i := range graph[cur] {
			if colors[i] == 0 {
				queue = append(queue, i)
				// 染成另一种颜色
				colors[i] = -colors[cur]
			} else if colors[i] == colors[cur] {
				return false
			}
		}
	}
	return true
}
func main() {
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	fmt.Println(isBipartite(graph))
}
