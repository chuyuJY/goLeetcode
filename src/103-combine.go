package main

import "fmt"

func combine(n int, k int) [][]int {
	res := [][]int{}
	var backTracking func(start, end, k int, combination []int)
	backTracking = func(start, end, k int, combination []int) {
		if k == 0 {
			temp := make([]int, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		if end-start+1 < k {
			return
		}
		for i := start; i <= end; i++ {
			combination[len(combination)-k] = i
			backTracking(i+1, end, k-1, combination)
		}
		return
	}
	combination := make([]int, k)
	backTracking(1, n, k, combination)
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
