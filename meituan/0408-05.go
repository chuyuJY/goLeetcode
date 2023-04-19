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
	ps := []int{0, 1}
	for i := 2; i <= n; i++ {
		var p int
		fmt.Fscan(sc, &p)
		ps = append(ps, p)
	}
	fmt.Fscanln(sc)
	var str string
	fmt.Fscanln(sc, &str)
	str = "#" + str
	test17(n, ps, str)
}

func test17(n int, ps []int, str string) {
	dp := make([][3]int, n+1) // [R, G, B]
	graph := make([][]int, n+1)
	for i := 2; i < len(ps); i++ {
		graph[ps[i]] = append(graph[ps[i]], i)
	}
	// 得到每个子树的色彩情况
	var dfsGetDp func(cur int)
	dfsGetDp = func(cur int) {
		if str[cur] == 'R' {
			dp[cur][0]++
		} else if str[cur] == 'G' {
			dp[cur][1]++
		} else if str[cur] == 'B' {
			dp[cur][2]++
		}

		for _, next := range graph[cur] {
			dfsGetDp(next)
			dp[cur][0] += dp[next][0]
			dp[cur][1] += dp[next][1]
			dp[cur][2] += dp[next][2]
		}
		return
	}
	// 寻找边
	res := 0
	var dfsGetEdge func(cur int)
	dfsGetEdge = func(cur int) {
		if isValid(dp[1], dp[cur]) {
			res++
		}
		for _, next := range graph[cur] {
			dfsGetEdge(next)
		}
	}
	dfsGetDp(1)
	dfsGetEdge(1)
	fmt.Println(res)
}

func isValid(sum, cur [3]int) bool {
	for i := 0; i < 3; i++ {
		if cur[i] == 0 || sum[i]-cur[i] == 0 {
			return false
		}
	}
	return true
}
