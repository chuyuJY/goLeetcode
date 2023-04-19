package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(sc, &n)
	ids, parents, vals := []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		var id, parent, val int
		fmt.Fscanln(sc, &id, &parent, &val)
		ids = append(ids, id)
		parents = append(parents, parent)
		vals = append(vals, val)
	}
	test2(ids, parents, vals)
}

func test2(ids, parents, vals []int) {
	start := -1
	nodeParent := make([]int, len(parents))
	nodeVals := make([]int, len(vals))
	graph := make([][]int, len(ids))
	for i := 0; i < len(ids); i++ {
		nodeVals[ids[i]] = vals[i]
		nodeParent[ids[i]] = parents[i]
		if parents[i] == -1 {
			start = ids[i]
			continue
		}
		graph[parents[i]] = append(graph[parents[i]], ids[i])
	}
	queue := []int{start}
	dp := make([]int, len(ids))
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if nodeParent[cur] == -1 {
			dp[cur] = nodeVals[cur]
		} else {
			dp[cur] = max(dp[nodeParent[cur]]+nodeVals[cur], nodeVals[cur])
		}
		for _, next := range graph[cur] {
			queue = append(queue, next)
		}
	}
	res := math.MinInt
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	fmt.Println(res)
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
