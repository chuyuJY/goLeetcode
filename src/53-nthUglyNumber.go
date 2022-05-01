package main

import "fmt"

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	a, b, c := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
		// 注意：此处不可以是if else结构，只要是相等发生，都必须++，因为要保证dp[a,b,c]*2,3,5都要大于dp[i-1]才可以
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}
	return dp[n]
}

//func min(a, b, c int) int {
//	if a < b {
//		if a < c {
//			return a
//		} else {
//			return c
//		}
//	} else {
//		if b < c {
//			return b
//		} else {
//			return c
//		}
//	}
//}
func main() {
	fmt.Println(nthUglyNumber(7))
}
