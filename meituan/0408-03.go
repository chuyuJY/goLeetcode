package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, m, s int
	fmt.Fscanln(sc, &n, &m, &s)
	nums := []int{}
	for i := 0; i < n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		nums = append(nums, ai)
	}
	fmt.Fscanln(sc)
	test15(n, m, s, nums)
}

func test15(n, m, s int, nums []int) {
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i < len(dp); i++ {
		u, v := nums[i-1], nums[i-1]
		for j := i; j >= 1 && j > i-m; j-- {
			u, v = max(u, nums[j-1]), min(v, nums[j-1])
			dp[i] = min(dp[i], dp[j-1]+(i-j+1)*((u+v)/2)+s)
		}
	}
	fmt.Println(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
