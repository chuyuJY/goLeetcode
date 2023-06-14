package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m, t int
	fmt.Fscanln(sc, &n, &m, &t)
	guns := [][3]int{}
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscanln(sc, &a, &b, &c)
		guns = append(guns, [3]int{a, b, c})
	}
	test3(m, t, guns)
}

// 一个大炮可以发射若干枚，并且不同大炮可以同时发射，若该条件是不同大炮不能同时发射，转换为二维完全背包问题，时间和弹药是消耗量。
// 完全背包
//func test3(m, t int, guns [][3]int) {
//	dp := make([][]int, m+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, t+1)
//	}
//	for _, gun := range guns {
//		for i := gun[1]; i < len(dp); i++ {
//			for j := gun[2]; j < len(dp[i]); j++ {
//				dp[i][j] = max(dp[i][j], dp[i-gun[1]][j-gun[2]]+gun[0])
//			}
//		}
//	}
//	fmt.Println(dp[m][t])
//}

// 因为能同时开炮，时间就只能对单个炮进行约束了，进而转换成多重背包(01背包)
func test3(m, t int, guns [][3]int) {
	dp := make([]int, m+1)
	for _, gun := range guns {
		for i := m; i >= gun[1]; i-- { // 01背包从后往前，完全背包从前往后
			times := min(i/gun[1], t/gun[2]) // 得到该炮的最大发射次数
			for j := 1; j <= times; j++ {
				dp[i] = max(dp[i], dp[i-j*gun[1]]+j*gun[0])
			}
		}
	}
	fmt.Println(dp[m])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
