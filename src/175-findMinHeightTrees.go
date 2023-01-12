package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 1. 建图和入度表
	graph, inDegrees := map[int][]int{}, map[int]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegrees[edge[0]]++
		inDegrees[edge[1]]++
	}
	// BFS
	res := []int{}
	queue := []int{}
	nextQueue := []int{}
	for node, inDegree := range inDegrees {
		if inDegree == 1 {
			queue = append(queue, node)
		}
	}
	count := len(queue)
	if n == count {
		res = append(res, queue...)
		return res
	}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, nextNode := range graph[curNode] {
			inDegrees[nextNode]--
			if inDegrees[nextNode] == 1 {
				nextQueue = append(nextQueue, nextNode)
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			if len(queue) == n-count {
				res = append(res, queue...)
				return res
			}
			count += len(queue)
			nextQueue = []int{}
		}
	}
	return res
}
