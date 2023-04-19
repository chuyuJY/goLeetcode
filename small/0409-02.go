package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, X, C int
	fmt.Fscanln(sc, &n, &X, &C)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		var xi int
		fmt.Fscan(sc, &xi)
		xs[i] = xi
	}
	fmt.Fscanln(sc)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		var yi int
		fmt.Fscan(sc, &yi)
		ys[i] = yi
	}
	fmt.Fscanln(sc)
	test5(X, C, xs, ys)
}

func test5(X, C int, xs, ys []int) {
	dp := make([]int, C+1)
	for j := 0; j < len(xs); j++ {
		dp[xs[j]] = max(dp[xs[j]], ys[j])
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < i; j++ {
			// 体积为 j 和 i-j 的方案都得存在才能继续
			if dp[j] == 0 || dp[i-j] == 0 {
				continue
			}
			if i-j == j {
				dp[i] = max(dp[i], dp[j]+dp[i-j]+X)
			} else {
				dp[i] = max(dp[i], dp[j]+dp[i-j])
			}
		}
	}
	fmt.Println(dp[C])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
