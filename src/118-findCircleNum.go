package main

// 1. 图搜索
//func findCircleNum(isConnected [][]int) int {
//	visited := make([]bool, len(isConnected))
//	res := 0
//	for node, visit := range visited {
//		if !visit {
//			bfsFindCircle(isConnected, node, visited)
//			res++
//		}
//	}
//	return res
//}
//
//func bfsFindCircle(isConnected [][]int, curNode int, visited []bool) {
//	queue := []int{curNode}
//	visited[curNode] = true
//	for len(queue) > 0 {
//		preNode := queue[0]
//		queue = queue[1:]
//		for node, connect := range isConnected[preNode] {
//			if connect == 1 && !visited[node] {
//				queue = append(queue, node)
//				visited[node] = true
//			}
//		}
//	}
//}
//////////////////////////////////////////////////////////////////////////

// 2. 并查集
func findCircleNum(isConnected [][]int) int {
	fathers := make([]int, len(isConnected))
	for node := range fathers {
		fathers[node] = node
	}
	res := len(isConnected)
	for nodeI := 0; nodeI < len(isConnected); nodeI++ {
		for nodeJ := nodeI + 1; nodeJ < len(isConnected[nodeI]); nodeJ++ {
			if isConnected[nodeI][nodeJ] == 1 && union(fathers, nodeI, nodeJ) {
				res--
			}
		}
	}
	return res
}

// 因为根节点的father还是自己，因此可定义寻找根节点的函数findFather
// 最终，路径上的节点都指向根节点，也即路径压缩
func findFather(fathers []int, node int) int {
	if fathers[node] != node {
		fathers[node] = findFather(fathers, fathers[node])
	}
	return fathers[node]
}

// 合并两个子集
// 其实就是将两个子集的根节点合到一起（把其中一个根节点的father改成另一个子集的根节点即可）
//func union(fathers []int, nodeI, nodeJ int) bool {
//	fatherI, fatherJ := findFather(fathers, nodeI), findFather(fathers, nodeJ)
//	// 若是不同的子集
//	if fatherI != fatherJ {
//		fathers[fatherI] = fatherJ
//		return true
//	}
//	// 若是相同的子集，无法合并
//	return false
//}
