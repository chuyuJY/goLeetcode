package main

// 1. 单向BFS
// func numBusesToDestination(routes [][]int, source int, target int) int {
// 	if source == target {
// 		return 0
// 	}
// 	// 1. 建图
// 	stationMap := map[int][]int{}              // 站点: []路线{}
// 	for routeIndex, stations := range routes { // routeIndex: 路线编号, stations: 站点
// 		for _, station := range stations {
// 			stationMap[station] = append(stationMap[station], routeIndex)
// 		}
// 	}
// 	// 2. BFS
// 	res := 1
// 	visited := map[int]bool{}
// 	queue := []int{}
// 	// 初始化第一轮
// 	for _, routeIndex := range stationMap[source] {
// 		visited[routeIndex] = true
// 		queue = append(queue, routes[routeIndex]...)
// 	}
// 	nextQueue := []int{}
// 	for len(queue) > 0 {
// 		curStation := queue[0]
// 		queue = queue[1:]
// 		if curStation == target {
// 			return res
// 		}
// 		for _, routeIndex := range stationMap[curStation] {
// 			if !visited[routeIndex] {
// 				nextQueue = append(nextQueue, routes[routeIndex]...) // 将未访问路线的所有站点加入
// 				visited[routeIndex] = true
// 			}
// 		}
// 		if len(queue) == 0 {
// 			queue = nextQueue
// 			nextQueue = []int{}
// 			res++
// 		}
// 	}
// 	return -1
// }

// 2. 双向BFS
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	// 1. 建图
	stationMap := map[int][]int{}
	for routeIndex, route := range routes {
		for _, station := range route {
			stationMap[station] = append(stationMap[station], routeIndex)
		}
	}
	// 2. BFS
	res := 1
	set1, set2 := map[int]bool{source: true}, map[int]bool{target: true} // station: bool
	visited := map[int]bool{}                                            // routeIndex: bool
	for len(set1) > 0 && len(set2) > 0 {
		if len(set1) > len(set2) {
			set1, set2 = set2, set1
		}
		set3 := map[int]bool{}
		for station := range set1 {
			for _, routeIndex := range stationMap[station] {
				if !visited[routeIndex] {
					for _, station := range routes[routeIndex] {
						if set2[station] {
							return res
						}
						set3[station] = true
					}
					visited[routeIndex] = true
				}
			}
		}
		set1 = set3
		res++
	}
	return -1
}
