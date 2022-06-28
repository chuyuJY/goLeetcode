package main

import "fmt"

// 1. 单向BFS
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	queue1, queue2 := []string{beginWord}, []string{}
//	visitedMap := map[string]bool{beginWord: true}
//	for _, word := range wordList {
//		visitedMap[word] = false
//	}
//	res := 1
//	for len(queue1) != 0 {
//		curWord := queue1[0]
//		queue1 = queue1[1:]
//		if curWord == endWord {
//			return res
//		}
//		neighbors := getNeighbors(curWord)
//		for _, neighbor := range neighbors {
//			if visited, ok := visitedMap[neighbor]; ok && !visited {
//				queue2 = append(queue2, neighbor)
//				visitedMap[neighbor] = true
//			}
//		}
//		if len(queue1) == 0 && len(queue2) != 0 {
//			queue1 = queue2
//			queue2 = []string{}
//			res++
//		}
//	}
//	return 0
//}

// 2. 双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	set1, set2 := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	visitedMap := map[string]bool{}
	for _, word := range wordList {
		visitedMap[word] = false
	}
	visitedMap[beginWord] = true
	// 若无该结尾单词
	if _, ok := visitedMap[endWord]; !ok {
		return 0
	}
	visitedMap[endWord] = true
	res := 2
	for len(set1) != 0 && len(set2) != 0 {
		// 寻找数量少的集合
		if len(set1) > len(set2) {
			set1, set2 = set2, set1
		}
		set3 := map[string]bool{}
		for curWord := range set1 {
			neighbors := getNeighbors(curWord)
			for _, neighbor := range neighbors {
				if set2[neighbor] {
					return res
				}
				if visited, ok := visitedMap[neighbor]; ok && !visited {
					set3[neighbor] = true
					visitedMap[neighbor] = true
				}
			}
		}
		set1 = set3
		res++
	}
	return 0
}

// 找出所有的可能, 只改变一个字母
func getNeighbors(curWord string) []string {
	neighbors := []string{}
	curBytes := []byte(curWord)
	for i, char := range curBytes {
		for newChar := 'a'; newChar <= 'z'; newChar++ {
			if byte(newChar) != char {
				curBytes[i] = byte(newChar)
				neighbors = append(neighbors, string(curBytes))
			}
		}
		curBytes[i] = char
	}
	return neighbors
}

func main() {
	wordList := []string{"hot", "dot", "dog", "lot", "log"}
	res := ladderLength("hit", "cog", wordList)
	fmt.Println(res)
}
