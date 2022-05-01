package main

import (
	"fmt"
)

//func cuttingRope1(n int) int {
//	x, y, temp, xRope := 0, 0, 1, 0
//	for i := 2; i <= n; i++ {
//		x = n / i
//		y = n % i
//		temp = 1
//		for j := 0; j < i; j++ {
//			if j < y {
//				temp *= (x + 1)
//			} else {
//				temp *= x
//			}
//		}
//		if temp > xRope {
//			xRope = temp
//		}
//	}
//	return xRope
//}

// 动态规划
func cuttingRope1(n int) int {
	// 用来记录每个长度的最优解
	dp := make([]int, n+1) // 0 1 其实没用到
	dp[2] = 1
	// 得到dp[3]-dp[n]的解
	for i := 3; i <= n; i++ {
		// 从第一段绳子=2开始（因为1的话，无增益）
		// 对于每一次来说，只需要比较三种情况即可
		for j := 2; j < i; j++ {
			x, y, z := j*dp[i-j], j*(i-j), dp[i]
			if x >= y && x >= z {
				dp[i] = x
			}
			if y >= x && y >= z {
				dp[i] = y
			}
			if z >= x && z >= y {
				dp[i] = z
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(cuttingRope1(3))
}
