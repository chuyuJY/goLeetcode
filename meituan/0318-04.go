package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, x, y int
	fmt.Fscanln(sc, &n, &x, &y)
	prices := [][]int{}
	for i := 0; i < n; i++ {
		var ori, dis int
		fmt.Fscanln(sc, &ori, &dis)
		prices = append(prices, []int{ori, dis})
	}
	test8(x, y, prices)
}

func test8(x, y int, prices [][]int) {
	dp := make([][]int, x+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, y+1)
	}
	for _, price := range prices {
		for i := x; i >= price[1]; i-- {
			for j := y; j >= 0; j-- {
				if i >= price[0] {
					dp[i][j] = max(dp[i][j], dp[i-price[0]][j]+1)
				}
				if i >= price[1] && j >= 1 {
					dp[i][j] = max(dp[i][j], dp[i-price[1]][j-1]+1)
				}
			}
		}
	}
	count, cost := 0, 0
	for i := x; i >= 0; i-- {
		if count < dp[i][y] {
			count = dp[i][y]
			cost = i
		} else if count == dp[i][y] {
			cost = min(cost, i)
		}
	}
	fmt.Println(count, cost)
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
