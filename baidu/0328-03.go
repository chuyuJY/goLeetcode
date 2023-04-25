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
	vals := []int{0}
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(sc, &v)
		vals = append(vals, v)
	}
	fmt.Fscanln(sc)
	graph := map[int][]int{}
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(sc, &u, &v)
		graph[u] = append(graph[u], v)
	}
	var q int
	fmt.Fscanln(sc, &q)
	opts := [][]int{}
	for i := 0; i < q; i++ {
		var t, g int
		fmt.Fscanln(sc, &t, &g)
		opts = append(opts, []int{t, g})
	}

	test5(graph, vals, opts)
}

func test5(graph map[int][]int, vals []int, opts [][]int) {
	// 1. 得到初始化状态各节点的 [2, 5]
	dp := make([][]int, len(vals))
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, 2)
		getTwoFive(vals[i], dp[i])
	}
	// 2. dfs 执行操作
	var dfsOpt func(cur int, val int)
	dfsOpt = func(cur int, val int) {
		getTwoFive(val, dp[cur])
		for _, next := range graph[cur] {
			dfsOpt(next, val)
		}
	}
	for _, opt := range opts {
		dfsOpt(opt[0], opt[1])
	}
	// 3. dfs 计算所有节点的0
	var dfsAdd func(cur int) (int, int)
	dfsAdd = func(cur int) (int, int) {
		for _, next := range graph[cur] {
			twoNext, fiveNext := dfsAdd(next)
			dp[cur][0] += twoNext
			dp[cur][1] += fiveNext
		}
		return dp[cur][0], dp[cur][1]
	}
	dfsAdd(1)
	// 4. 输出结果
	for i := 1; i < len(dp); i++ {
		if i != len(dp)-1 {
			fmt.Printf("%v ", min(dp[i][0], dp[i][1]))
		} else {
			fmt.Printf("%v", min(dp[i][0], dp[i][1]))
		}
	}
}

func getTwoFive(val int, nums []int) {
	two, five := 0, 0
	for val > 0 && val%2 == 0 {
		two++
		val /= 2
	}
	for val > 0 && val%5 == 0 {
		five++
		val /= 5
	}
	nums[0] += two
	nums[1] += five
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
