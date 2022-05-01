package main

import (
	"fmt"
	"strconv"
)

// 动态规划
//func translateNum(num int) int {
//	// 排除 输入0 这种情况
//	if num == 0 {
//		return 1
//	}
//	// 求数字位数
//	temp, cnt := num, 0
//	for temp > 0 {
//		temp /= 10
//		cnt++
//	}
//	// 设置dp
//	dp := make([]int, cnt+1)
//	dp[0], dp[1] = 1, 1
//	for i := 2; i <= cnt; i++ {
//		// 求i所在位和前一数字是否能满足条件
//		// 1.这是采用取余和整除法求两位数字
//		temp = num % (int(math.Pow(10, float64(cnt-i+2)))) / (int(math.Pow(10, float64(cnt-i))))
//		// 2.也可以用转换成字符串的形式求
//
//		if temp >= 10 && temp <= 25 {
//			// 转移方程
//			dp[i] = dp[i-1] + dp[i-2]
//		} else {
//			dp[i] = dp[i-1]
//		}
//	}
//	return dp[cnt]
//}

// 递归YYDS
func translateNum(num int) int {
	str := strconv.Itoa(num)
	return translateStr(str, len(str)-1)
}
func translateStr(str string, i int) int {
	if i < 1 {
		return 1
	}
	temp := string(str[i-1]) + string(str[i])
	if temp >= "10" && temp <= "25" {
		return translateStr(str, i-1) + translateStr(str, i-2)
	} else {
		return translateStr(str, i-1)
	}
}
func main() {
	fmt.Println(translateNum(12258))
}
