package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// 处理空矩阵情况
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 此时矩阵不为空
	// 定义上下左右四个边界，对边界按顺序进行遍历就完事了
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := make([]int, (bottom+1)*(right+1))
	cnt := 0
	for {
		for i := left; i <= right; i++ {
			res[cnt] = matrix[top][i]
			cnt++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[cnt] = matrix[i][right]
			cnt++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res[cnt] = matrix[bottom][i]
			cnt++
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res[cnt] = matrix[i][left]
			cnt++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
func main() {
	//matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	matrix := [][]int{{7}, {9}, {6}}
	fmt.Println(spiralOrder(matrix))
}
