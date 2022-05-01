package main

import "fmt"

//func numWays(n int) int {
//	if n == 0 || n == 1 {
//		return 1
//	}
//	var (
//		a = 1
//		b = 1
//	)
//	for i := 2; i <= n; i++ {
//		a, b = b, (a+b)%(1e9+7)
//	}
//	return b
//}

var nums []int

func numWays(n int) int {
	nums = make([]int, n+1)
	Ways(n)
	return nums[n]
}
func Ways(n int) int {
	if n < 2 {
		nums[n] = 1
		return nums[n]
	} else {
		if nums[n] == 0 {
			nums[n] = (Ways(n-1) + Ways(n-2)) % 1000000007
			return nums[n]
		} else {
			return nums[n]
		}
	}
}

func main() {
	fmt.Println(numWays(7))
}
