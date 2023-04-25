package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(sc, &n, &k)
	str, _ := sc.ReadString('\n')
	color := strings.TrimRight(str, "\r\n")
	graph := map[int][]int{}
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(sc, &u, &v)
		graph[u] = append(graph[u], v)
	}
	color = "#" + color
	test8(color, graph, k)
}

func test8(color string, graph map[int][]int, k int) {
	nums := []int{}
	var dfs func(cur int, flag bool) int
	dfs = func(cur int, flag bool) int {
		curRes := 0
		curFlag := false
		if color[cur] != 'R' {
			curFlag = true
		} else {
			curRes++
		}
		for _, next := range graph[cur] {
			curRes += dfs(next, curFlag)
		}
		if flag {
			nums = append(nums, curRes)
		}
		if color[cur] != 'R' {
			curRes = 0
		}
		return curRes
	}
	if color[1] == 'R' {
		dfs(1, true)
	} else {
		dfs(1, false)
	}
	if len(nums) < k {
		fmt.Println(-1)
		return
	}
	sort.Ints(nums)
	if nums[len(nums)-k] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(nums[len(nums)-k])
	}
}
