package main

import (
	"fmt"
	"time"
)

//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[i]); j++ {
//			if matrix[i][j] == target {
//				return true
//			}
//		}
//	}
//	return false
//}

// 从右上角开始判断！
//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	if len(matrix) == 0 {
//		return false
//	}
//	j := len(matrix[0]) - 1
//	for i := 0; i < len(matrix) && j >= 0; { // j >= 0 很关键 不仅可以防止 matrix = [[]]的情况，也可以用作循环终止条件
//		if matrix[i][j] < target {
//			i++
//		} else if matrix[i][j] > target {
//			j--
//		} else {
//			return true
//		}
//	}
//	return false
//}

// 从左下角开始判断！
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{}
	target := 30
	initTime := time.Now()
	in := findNumberIn2DArray(matrix, target)
	fmt.Println(in)
	useTime := time.Since(initTime)
	fmt.Printf("%s", useTime)
}
