package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(sc, &n, &m)
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		fmt.Fscan(sc, &nums[i])
	}
	fmt.Fscanln(sc)
	test2(nums, m)
}

func test2(nums []int, m int) {
	dp := make([]int, m+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = m + 1
	}
	for _, num := range nums {
		for i := m; i >= num; i-- {
			dp[i] = min(dp[i], dp[i-num]+1)
		}
	}
	if dp[m] == m+1 {
		fmt.Println("No solution")
	} else {
		fmt.Println(dp[m])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
