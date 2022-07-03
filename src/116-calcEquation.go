package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, len(queries))
	graph := buildGraph(equations, values)
	for index, query := range queries {
		from := query[0]
		to := query[1]
		if _, ok := graph[from]; !ok {
			res[index] = -1
		} else {
			// 搜索路径
			visited := map[string]bool{}
			// 1. DFS
			//res[index] = dfs(from, to, graph, visited)
			// 2. BFS
			res[index] = bfs(from, to, graph, visited)
		}
	}
	return res
}

// 1. 构建图
func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	graph := map[string]map[string]float64{}
	for index, equation := range equations {
		node1 := equation[0]
		node2 := equation[1]
		if graph[node1] == nil {
			graph[node1] = map[string]float64{}
		}
		if graph[node2] == nil {
			graph[node2] = map[string]float64{}
		}
		graph[node1][node2] = values[index]
		graph[node2][node1] = 1.0 / values[index]
	}
	return graph
}

// 2. 计算路径dfs
//func dfs(from, to string, graph map[string]map[string]float64, visited map[string]bool) float64 {
//	if from == to {
//		return 1.0
//	}
//	visited[from] = true
//	if nodes, ok := graph[from]; ok {
//		for newFrom, value := range nodes {
//			if !visited[newFrom] {
//				nextValue := dfs(newFrom, to, graph, visited)
//				if nextValue > 0 {
//					return value * nextValue
//				}
//			}
//		}
//	}
//	delete(visited, from)
//	return -1.0
//}

// 3. 计算路径bfs  // todo
//func bfs(from, to string, graph map[string]map[string]float64, visited map[string]bool) float64 {
//
//}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}
