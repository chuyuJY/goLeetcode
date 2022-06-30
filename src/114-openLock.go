package main

import "fmt"

// 1. 单向BFS
func openLock(deadends []string, target string) int {
	beginWord := "0000"
	deadMap := map[string]bool{}
	for _, word := range deadends {
		deadMap[word] = true
	}
	if deadMap[beginWord] {
		return -1
	}
	// 广度优先遍历BFS
	res := 0
	queue1, queue2 := []string{beginWord}, []string{}
	visitedMap := map[string]bool{beginWord: true}
	for len(queue1) != 0 {
		curWord := queue1[0]
		queue1 = queue1[1:]
		if curWord == target {
			return res
		}
		neighbors := getNeighbors(curWord)
		for _, neighbor := range neighbors {
			if !deadMap[neighbor] && !visitedMap[neighbor] {
				queue2 = append(queue2, neighbor)
				visitedMap[neighbor] = true
			}
		}
		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = []string{}
			res++
		}
	}
	return -1
}

// 2. 单向BFS
//func openLock(deadends []string, target string) int {
//	beginWord := "0000"
//	if beginWord == target {
//		return 0
//	}
//	set1, set2 := map[string]bool{beginWord: true}, map[string]bool{target: true}
//	visitedMap := map[string]bool{beginWord: true, target: true}
//	deadMap := map[string]bool{}
//	for _, word := range deadends {
//		deadMap[word] = true
//	}
//	if deadMap[beginWord] {
//		return -1
//	}
//	res := 1
//	for len(set1) != 0 && len(set2) != 0 {
//		if len(set1) > len(set2) {
//			set1, set2 = set2, set1
//		}
//		set3 := map[string]bool{}
//		for curWord := range set1 {
//			neighbors := getNeighbors(curWord)
//			for _, neighbor := range neighbors {
//				if set2[neighbor] {
//					return res
//				}
//				if !visitedMap[neighbor] && !deadMap[neighbor] {
//					set3[neighbor] = true
//					visitedMap[neighbor] = true
//				}
//			}
//		}
//		set1 = set3
//		res++
//	}
//	return -1
//}

//func getNeighbors(curWord string) []string {
//	neighbors := []string{}
//	curBytes := []byte(curWord)
//	for i, char := range curBytes {
//		if char == '9' {
//			curBytes[i] = '0'
//			neighbors = append(neighbors, string(curBytes))
//		} else {
//			curBytes[i] = char + 1
//			neighbors = append(neighbors, string(curBytes))
//		}
//		if char == '0' {
//			curBytes[i] = '9'
//			neighbors = append(neighbors, string(curBytes))
//		} else {
//			curBytes[i] = char - 1
//			neighbors = append(neighbors, string(curBytes))
//		}
//		curBytes[i] = char
//	}
//	return neighbors
//}

func main() {
	str := '2'
	fmt.Println(int(str))
}
