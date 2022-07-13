package main

import "fmt"

//// 1. 并查集
//func longestConsecutive(nums []int) int {
//	// 1. 构建子树和计数
//	fathers, count := map[int]int{}, map[int]int{}
//	for _, num := range nums {
//		fathers[num] = num
//		count[num] = 1
//	}
//	// 2. 合并子树
//	for _, num := range nums {
//		if _, exist := count[num-1]; exist {
//			union(fathers, count, num-1, num)
//		}
//		if _, exist := count[num+1]; exist {
//			union(fathers, count, num+1, num)
//		}
//	}
//	res := 0
//	for _, value := range count {
//		res = max(res, value)
//	}
//	return res
//}
//
//func findFather(fathers map[int]int, node int) int {
//	if fathers[node] != node {
//		fathers[node] = findFather(fathers, fathers[node])
//	}
//	return fathers[node]
//}
//
//func union(fathers, count map[int]int, nodeI, nodeJ int) {
//	fatherI, fatherJ := findFather(fathers, nodeI), findFather(fathers, nodeJ)
//	if fatherI != fatherJ {
//		fathers[fatherI] = fatherJ
//		count[fatherJ] += count[fatherI]
//	}
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2. 图搜索
func longestConsecutive(nums []int) int {
	graph := map[int]int{}
	for _, num := range nums {
		graph[num] = num
	}
	res := 0
	for curNode := range graph {
		res = max(res, bfsConsecutive(graph, curNode))
	}

	return res
}

func bfsConsecutive(graph map[int]int, curNode int) int {
	queue := []int{curNode}
	neighbors := []int{-1, 1}
	length := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		delete(graph, node)
		length++
		for _, neighbor := range neighbors {
			newNode := node + neighbor
			if _, exist := graph[newNode]; exist {
				queue = append(queue, graph[newNode])
			}
		}
	}
	return length
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
