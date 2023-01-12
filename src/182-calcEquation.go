package main

// 1. 并查集
// type node struct {
// 	value  float64
// 	parent *node
// }

// func newNode(value float64) *node {
// 	node := &node{
// 		value: value,
// 	}
// 	node.parent = node
// 	return node
// }

// // 返回 root
// func findRoot(node *node) *node {
// 	if node.parent == node {
// 		return node
// 	}
// 	node.parent = findRoot(node.parent) // 路径压缩
// 	return node.parent
// }

// // 合并子集
// // nodeI, nodeJ 是当前边的两个顶点, value 是边的值: nodeI/nodeJ, hashMap记录所有节点的值
// func union(nodeI, nodeJ *node, value float64, hashMap map[string]*node) {
// 	rootI, rootJ := findRoot(nodeI), findRoot(nodeJ)
// 	// 合并操作
// 	if rootI != rootJ {
// 		ratio := nodeJ.value * value / nodeI.value
// 		for k, v := range hashMap {
// 			if findRoot(v) == rootI {
// 				hashMap[k].value = v.value * ratio
// 			}
// 		}
// 		rootI.parent = rootJ
// 	}
// }

// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
// 	res := []float64{}
// 	hashMap := map[string]*node{}

// 	// 构建并查集
// 	for index, equation := range equations {
// 		nodeI, okI := hashMap[equation[0]]
// 		nodeJ, okJ := hashMap[equation[1]]
// 		// 1. 两个都没出现过
// 		if !okI && !okJ {
// 			nodeI, nodeJ = newNode(values[index]), newNode(1.0)
// 			nodeI.parent = nodeJ.parent
// 			hashMap[equation[0]] = nodeI
// 			hashMap[equation[1]] = nodeJ
// 		} else if !okI { // 2. 仅 nodeI 没出现过
// 			nodeI := newNode(nodeJ.value * values[index])
// 			nodeI.parent = nodeJ
// 			hashMap[equation[0]] = nodeI
// 		} else if !okJ { // 3. 仅 nodeJ 没出现过
// 			nodeJ := newNode(nodeI.value / values[index])
// 			nodeJ.parent = nodeI
// 			hashMap[equation[1]] = nodeJ
// 		} else { // 4. 都出现过
// 			union(nodeI, nodeJ, values[index], hashMap)
// 		}
// 	}

// 	// 计算等式
// 	for _, query := range queries {
// 		nodeI, okI := hashMap[query[0]]
// 		nodeJ, okJ := hashMap[query[1]]
// 		if okI && okJ && findRoot(nodeI) == findRoot(nodeJ) {
// 			res = append(res, nodeI.value/nodeJ.value)
// 		} else {
// 			res = append(res, -1.0)
// 		}
// 	}
// 	return res
// }

// 2. 图搜索

func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	graph := map[string]map[string]float64{}
	for index, equation := range equations {
		node1, node2 := equation[0], equation[1]
		graph[node1][node2] = values[index]
		graph[node2][node1] = 1.0 / values[index]
	}
}
