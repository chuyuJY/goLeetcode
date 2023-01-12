package main

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res, path := [][]string{}, []string{beginWord}
	visitedMap := map[string]bool{}
	for _, word := range wordList {
		visitedMap[word] = false
	}
	visitedMap[beginWord] = true
	// 如果endWord不在hash表中，表示不存在转换列表，返回空集合
	if _, exist := visitedMap[endWord]; !exist {
		return [][]string{}
	}
	// 建图
	graph := buildGraph(visitedMap)
	// 寻找endword到图中节点的最短路径长度
	distances := bfsGetDistances(endWord, graph)

	// dfs寻找最短路径
	var dfsGetPaths func(curWord string)
	dfsGetPaths = func(curWord string) {
		if curWord == endWord {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, path)
			return
		}
		for _, nextWord := range graph[curWord] {
			if len(path)+1 <= distances[beginWord]-distances[nextWord] {
				path = append(path, nextWord)
				visitedMap[nextWord] = true
				dfsGetPaths(nextWord)
				visitedMap[nextWord] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfsGetPaths(beginWord)
	return res
}

// 寻找endword到图中每个点的最短路径
func bfsGetDistances(endWord string, graph map[string][]string) map[string]int {
	distances := map[string]int{}
	distance := 1
	visitedMap := map[string]bool{endWord: true}
	queue := []string{endWord}
	nextQueue := []string{}
	for len(queue) > 0 {
		curWord := queue[0]
		queue = queue[1:]
		for _, nextWord := range graph[curWord] {
			if !visitedMap[nextWord] {
				distances[nextWord] = distance
				nextQueue = append(nextQueue, nextWord)
				visitedMap[nextWord] = true
			}
		}
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = []string{}
			distance++
		}
	}
	return distances
}

// 构建邻接表
func buildGraph(visitedMap map[string]bool) map[string][]string {
	graph := map[string][]string{}
	for curWord, _ := range visitedMap {
		curBytes := []byte(curWord)
		for index, curByte := range curBytes {
			for newByte := 'a'; newByte <= 'z'; newByte++ {
				if curByte != byte(newByte) {
					curBytes[index] = byte(newByte)
					if _, exist := visitedMap[string(curBytes)]; exist {
						graph[curWord] = append(graph[curWord], string(curBytes))
					}
				}
			}
			curBytes[index] = curByte
		}
	}
	return graph
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	findLadders(beginWord, endWord, wordList)
}
