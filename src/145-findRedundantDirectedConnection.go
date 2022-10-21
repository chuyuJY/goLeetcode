package main

import "fmt"

func findRedundantDirectedConnection(edges [][]int) []int {
	roots := make([]int, len(edges)+1)
	for node := 1; node < len(roots); node++ {
		roots[node] = node
	}
	twoDegreeNode, inDegree := 0, make([]int, len(edges)+1)
	for _, edge := range edges {
		inDegree[edge[1]]++
		if inDegree[edge[1]] == 2 {
			twoDegreeNode = edge[1]
		}
	}
	// 倒序，因为返回最后出现在数组中的边
	for index := len(edges) - 1; index >= 0; index-- {
		// 1. 没有入度为2的节点
		if twoDegreeNode == 0 {
			if !isCycle(roots, edges, index) {
				return edges[index]
			}
		} else if edges[index][1] == twoDegreeNode { // 2. 入度为2的节点
			if !isCycle(roots, edges, index) {
				return edges[index]
			}
		}
	}
	return nil
}

func isCycle(roots []int, edges [][]int, delEdge int) bool {
	// 重置roots
	for node := 1; node < len(roots); node++ {
		roots[node] = node
	}
	for index, edge := range edges {
		if index != delEdge && !union(roots, edge[0], edge[1]) {
			return true
		}
	}
	return false
}

func findRoot(roots []int, node int) int {
	if node != roots[node] {
		roots[node] = findRoot(roots, roots[node])
	}
	return roots[node]
}

func union(roots []int, nodeI, nodeJ int) bool {
	rootI, rootJ := findRoot(roots, nodeI), findRoot(roots, nodeJ)
	// 若是不同的子集
	if rootI != rootJ {
		roots[rootJ] = rootI
		return true
	}
	// 若是相同的子集，无法合并
	return false
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	fmt.Println(findRedundantDirectedConnection(edges))
}
