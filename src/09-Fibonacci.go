package main

import "fmt"

// 1. 递归法
//func fib(n int) int {
//	return fibonacci(n) % 1000000007
//}
//func fibonacci(n int) int {
//	if n > 1 {
//		return fibonacci(n-1) + fibonacci(n-2)
//	} else if n == 1 {
//		return 1
//	} else {
//		return 0
//	}
//}

// 2. 记忆递归法
// 定义一个全局数组，存储每次递归出来的值

//var fibNums []int
//func fib(n int) int {
//	fibNums = make([]int, n+1) // 注意是 n+1 个空间
//	return fibonacci(n)
//}
//func fibonacci(n int) int {
//	if n < 2 {
//		return n
//	} else {
//		if fibNums[n] == 0 {
//			fibNums[n] = (fibonacci(n-2) + fibonacci(n-1)) % 1000000007
//			return fibNums[n]
//		} else {
//			return fibNums[n]
//		}
//	}
//}

// 2. 动态规划
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	var (
		a = 0
		b = 1
		//temp int
	)
	for i := 2; i <= n; i++ {
		//temp = a
		//a = b
		//b = (temp + a) % 1000000007
		a, b = b, (b+a)%1000000007
	}
	return b
}

func main() {
	fmt.Println(fib(95))
}
