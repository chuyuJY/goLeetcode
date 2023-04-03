package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	fmt.Scanf("%v %v\n", &n, &m)
	sc := bufio.NewScanner(os.Stdin)
	matrix := []string{}
	if sc.Scan() {
		str := sc.Text()
		matrix = append(matrix, str)
	}
	fmt.Println(count(matrix))
}

//func count(matrix []string) int {
//	green, blue, both := 0, 0, 0
//	visitedMap := make([][]bool, len(matrix))
//	for i := 0; i < len(visitedMap); i++ {
//		visitedMap[i] = make([]bool, len(matrix[0]))
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			if !visitedMap[i][j] && matrix[i][j] == 'G' {
//				green++
//				bfs(matrix, 0, &visitedMap)
//			}
//		}
//	}
//	for i := 0; i < len(visitedMap); i++ {
//		for j := 0; j < len(visitedMap[i]); j++ {
//			visitedMap[i][j] = false
//		}
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			if !visitedMap[i][j] && matrix[i][j] == 'B' {
//				blue++
//				bfs(matrix, 1, &visitedMap)
//			}
//		}
//	}
//	for i := 0; i < len(visitedMap); i++ {
//		for j := 0; j < len(visitedMap[i]); j++ {
//			visitedMap[i][j] = false
//		}
//	}
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			if !visitedMap[i][j] && matrix[i][j] == 'A' {
//				both++
//				bfs(matrix, 2, &visitedMap)
//			}
//		}
//	}
//	return green + blue - both
//}

func bfs(matrix []string, k int, visitedMap *[][]bool) {
	// 0, 1, 2
	color := byte('R')
	if k == 0 {
		color = 'G'
	} else if k == 1 {
		color = 'B'
	} else if k == 2 {
		color = 'A'
	}
	queue := [][]int{{0, 0}}
	nextQueue := [][]int{}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[:len(queue)-1]
		for _, dir := range dirs {
			newX, newY := curNode[0]+dir[0], curNode[1]+dir[1]
			if 0 <= newX && newX < len(matrix) && 0 <= newY && newY < len(matrix[0]) && !visitedMap[newX][newY] {
				if color != 'A' && matrix[newX][newY] == color {
					nextQueue = append(nextQueue, []int{newX, newY})
					(*visitedMap)[newX][newY] = true
				} else if color == 'A' && matrix[newX][newY] == 'G' || matrix[newX][newY] == 'B' {
					nextQueue = append(nextQueue, []int{newX, newY})
					(*visitedMap)[newX][newY] = true
				}
			}
		}
		if len(queue) == 0 && len(nextQueue) > 0 {
			queue = nextQueue
			nextQueue = [][]int{}
		}
	}
}
