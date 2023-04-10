package main

func main() {

}

func test2(n, x, y int, edges [][]int) int {
	inDegrees := make([]int, n)
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegrees[edge[1]]++
	}
	queue := []int{}
	nextQueue := []int{}
	for i, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, i)
			break
		}
	}
	res := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range graph[cur] {
			if cur == x || cur == y {

			}
		}
	}
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
