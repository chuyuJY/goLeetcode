package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var mod int64 = 1e9 + 7

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	nums := []int{}
	for i := 0; i < n; i++ {
		var ai int
		fmt.Fscan(sc, &ai)
		nums = append(nums, ai)
	}
	fmt.Fscanln(sc)
	test7(nums)
}

func test7(nums []int) {
	var res int64 = 0
	sort.Ints(nums)
	dp := make([]int64, nums[len(nums)-1]+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		dp[i] = (1 + int64(i+1)*dp[i-1]) % mod
	}
	for i := 0; i < len(nums); i++ {
		res += dp[nums[i]] % mod
	}
	fmt.Println(res)
}
