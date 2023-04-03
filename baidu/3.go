package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewScanner(os.Stdin)
	nums := []int{}
	if sc.Scan() {
		str := sc.Text()
		strs := strings.Split(str, " ")
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
	}
	edges := [][]int{}
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Scanf("%v %v\n", &u, &v)
		edges = append(edges, []int{u, v})
	}
	var q int
	fmt.Scanf("%v\n", &q)
	ops := [][]int{}
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Scanf("%v %v\n", &x, &y)
		ops = append(ops, []int{x, y})
	}
	fmt.Println(test3(nums, edges, ops))
}

func test3(nums []int, edges [][]int, ops [][]int) int {
	graph := buildGraph(edges)
	visitedMap := map[int]bool{}
	var dfs func(node int, val int)
	dfs = func(node int, val int) {
		nextNodes := graph[node]
		for _, node := range nextNodes {
			if !visitedMap[node] {
				visitedMap[node] = true
				nums[node] *= val
				dfs(node, val)
			}
		}
	}
	for _, op := range ops {
		visitedMap = map[int]bool{op[0]: true}
		dfs(op[0], op[1])
	}
	res := []int{}

}

func buildGraph(edges [][]int) map[int][]int {
	graph := map[int][]int{}
	for _, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = []int{}
		}
		if graph[edge[1]] == nil {
			graph[edge[1]] = []int{}
		}
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

