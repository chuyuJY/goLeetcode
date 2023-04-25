package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	acks := [][]int{}
	for i := 0; i < n; i++ {
		var weight, skip int
		fmt.Fscanln(sc, &weight, &skip)
		acks = append(acks, []int{weight, skip})
	}
	test1(acks)
}

func test1(acks [][]int) {
	dp := make([]int, len(acks)+1)
	// dp[i] 表示 [i:] 最大结果
	for i := len(dp) - 2; i >= 0; i-- {
		next := i + acks[i][1] + 1
		if next < len(dp) {
			dp[i] = max(dp[i+1], dp[next]+acks[i][0])
		} else {
			dp[i] = max(dp[i+1], acks[i][0])
		}
	}
	fmt.Println(dp[0])
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
