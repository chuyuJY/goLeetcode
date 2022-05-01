package main

import "fmt"

// 1. 不考虑大数时，相当简单
//func printNumbers(n int) []int {
//	temp := 1
//	for i := 0; i < n; i++ {
//		temp *= 10
//	}
//	nums := make([]int, temp-1)
//	for i := 0; i < temp-1; i++ {
//		nums[i] = i + 1
//	}
//	return nums
//}
// 2. 大数，采用字符串模拟数字

func main() {
	fmt.Println(printNumbers(3))
}
